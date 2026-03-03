// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.30;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IValueStore} from "../interfaces/IValueStore.sol";

/// @title DIAOracleV3MetaFairValueField
/// @notice Aggregates values from multiple ValueStore contracts using median calculation
/// @author DIA (diadata.org)
/// @dev Returns median values from multiple oracle sources with timeout and threshold protection
contract DIAOracleV3MetaFairValueField is Ownable {
    struct MedianSet {
        uint256 fairValue;
        uint256 usdValue;
        uint256 numerator;
        uint256 denominator;
        uint256 timestamp;
    }

    /// @notice Mapping of value store addresses by index
    mapping(uint256 => address) public valueStores;
    /// @notice Number of registered value stores
    uint256 public numValueStores;
    /// @notice Minimum number of valid responses required
    uint256 public threshold;
    /// @notice Maximum age of data in seconds
    uint256 public timeoutSeconds;
    /// @notice Maximum number of value stores that can be registered
    uint256 public maxValueStores;

    bytes32 private constant _FAIR_VALUE = keccak256("fairValue");
    bytes32 private constant _USD_VALUE = keccak256("usdValue");
    bytes32 private constant _NUMERATOR = keccak256("numerator");
    bytes32 private constant _DENOMINATOR = keccak256("denominator");
    uint256 private constant _MAX_TIMEOUT_SECONDS = 1 days;
    uint256 private constant _DEFAULT_MAX_VALUE_STORES = 100;

    error InvalidThreshold(uint256);
    error InvalidTimeOut(uint256);
    error TimeoutExceedsLimit(uint256);
    error ThresholdNotMet(uint256, uint256);
    error ZeroAddress();
    error OracleExists();
    error OracleNotFound();
    error MaxValueStoresReached(uint256);

    /// @notice Emitted when timeout seconds is changed
    /// @param oldTimeoutSeconds The previous timeout value
    /// @param newTimeoutSeconds The new timeout value
    event TimeoutSecondsChanged(uint256 indexed oldTimeoutSeconds, uint256 indexed newTimeoutSeconds);

    /// @notice Emitted when threshold is changed
    /// @param oldThreshold The previous threshold value
    /// @param newThreshold The new threshold value
    event ThresholdChanged(uint256 indexed oldThreshold, uint256 indexed newThreshold);

    /// @notice Emitted when max value stores is changed
    /// @param oldMaxValueStores The previous max value stores
    /// @param newMaxValueStores The new max value stores
    event MaxValueStoresChanged(uint256 indexed oldMaxValueStores, uint256 indexed newMaxValueStores);

    /// @notice Emitted when a value store is added
    /// @param storeAddress The address of the added store
    /// @param storeIndex The index of the added store
    event ValueStoreAdded(address indexed storeAddress, uint256 indexed storeIndex);

    /// @notice Emitted when a value store is removed
    /// @param storeAddress The address of the removed store
    /// @param storeIndex The index of the removed store
    event ValueStoreRemoved(address indexed storeAddress, uint256 indexed storeIndex);

    /// @notice Constructor to initialize the oracle with an owner
    /// @param initialOwner The address that will own the contract
    constructor(address initialOwner) Ownable(initialOwner) {
        if (initialOwner == address(0)) revert ZeroAddress();
        maxValueStores = _DEFAULT_MAX_VALUE_STORES;
    }

    /// @notice Add a new ValueStore to the oracle
    /// @param newStore The address of the ValueStore to add
    function addValueStore(address newStore) external onlyOwner {
        if (newStore == address(0)) revert ZeroAddress();
        if (numValueStores >= maxValueStores) revert MaxValueStoresReached(maxValueStores);
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == newStore) revert OracleExists();
        }
        valueStores[numValueStores++] = newStore;
        emit ValueStoreAdded(newStore, numValueStores - 1);
    }

    /// @notice Remove a ValueStore from the oracle
    /// @param storeAddr The address of the ValueStore to remove
    function removeValueStore(address storeAddr) external onlyOwner {
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == storeAddr) {
                valueStores[i] = valueStores[numValueStores - 1];
                delete valueStores[numValueStores - 1];
                numValueStores--;
                emit ValueStoreRemoved(storeAddr, i);
                return;
            }
        }
        revert OracleNotFound();
    }

    /// @notice Set the minimum number of valid responses required
    /// @param newThreshold The new threshold value
    function setThreshold(uint256 newThreshold) external onlyOwner {
        if (newThreshold == 0) revert InvalidThreshold(newThreshold);
        uint256 oldThreshold = threshold;
        threshold = newThreshold;
        emit ThresholdChanged(oldThreshold, newThreshold);
    }

    /// @notice Set the maximum age of data considered valid
    /// @param newTimeoutSeconds The new timeout in seconds
    function setTimeoutSeconds(uint256 newTimeoutSeconds) external onlyOwner {
        if (newTimeoutSeconds == 0) {
            revert InvalidTimeOut(newTimeoutSeconds);
        }
        if (newTimeoutSeconds > _MAX_TIMEOUT_SECONDS) {
            revert TimeoutExceedsLimit(newTimeoutSeconds);
        }
        uint256 oldTimeoutSeconds = timeoutSeconds;
        timeoutSeconds = newTimeoutSeconds;
        emit TimeoutSecondsChanged(oldTimeoutSeconds, newTimeoutSeconds);
    }

    /// @notice Set the maximum number of value stores that can be registered
    /// @param newMaxValueStores The new maximum value stores
    function setMaxValueStores(uint256 newMaxValueStores) external onlyOwner {
        if (newMaxValueStores == 0) revert InvalidThreshold(newMaxValueStores);
        if (newMaxValueStores < numValueStores) {
            revert InvalidThreshold(newMaxValueStores); // Cannot set below current count
        }
        uint256 oldMaxValueStores = maxValueStores;
        maxValueStores = newMaxValueStores;
        emit MaxValueStoresChanged(oldMaxValueStores, newMaxValueStores);
    }

    /// @notice Get median values from all registered ValueStores
    /// @param key The key to query from ValueStores
    /// @return median The median of all valid values
    function getMedianValues(string memory key) public view returns (MedianSet memory median) {
        (
            uint256[] memory fairValues,
            uint256[] memory usdValues,
            uint256[] memory nums,
            uint256[] memory dens,
            uint256[] memory timestamps
        ) = _initializeValueArrays();

        uint256 count = _collectValues(key, fairValues, usdValues, nums, dens, timestamps);

        _ensureThresholdMet(count);

        _sortValues(fairValues, usdValues, nums, dens, timestamps, count);

        return _calculateMedian(fairValues, usdValues, nums, dens, timestamps, count);
    }

    function _initializeValueArrays()
        private
        view
        returns (
            uint256[] memory fairValues,
            uint256[] memory usdValues,
            uint256[] memory nums,
            uint256[] memory dens,
            uint256[] memory timestamps
        )
    {
        fairValues = new uint256[](numValueStores);
        usdValues = new uint256[](numValueStores);
        nums = new uint256[](numValueStores);
        dens = new uint256[](numValueStores);
        timestamps = new uint256[](numValueStores);
    }

    function _collectValues(
        string memory key,
        uint256[] memory fairValues,
        uint256[] memory usdValues,
        uint256[] memory nums,
        uint256[] memory dens,
        uint256[] memory timestamps
    ) private view returns (uint256 count) {
        uint256 storeCount = numValueStores;
        for (uint256 i = 0; i < storeCount; ++i) {
            IValueStore store = IValueStore(valueStores[i]);
            try store.getValue(key) returns (uint256 fairV, uint256 usdV, uint256 num, uint256 den, uint256 ts) {
                if (ts + timeoutSeconds < block.timestamp) continue;

                fairValues[count] = fairV;
                usdValues[count] = usdV;
                nums[count] = num;
                dens[count] = den;
                timestamps[count] = ts;
                count++;
            } catch {
                continue;
            }
        }
    }

    function _ensureThresholdMet(uint256 count) private view {
        if (count < threshold) revert ThresholdNotMet(count, threshold);
    }

    function _sortValues(
        uint256[] memory fairValues,
        uint256[] memory usdValues,
        uint256[] memory nums,
        uint256[] memory dens,
        uint256[] memory timestamps,
        uint256 count
    ) private pure {
        if (count == 0) return;

        uint256 fairSum = 0;
        for (uint256 i = 0; i < count; ++i) {
            fairSum += fairValues[i];
        }

        if (fairSum != 0) {
            _sortMultipleByReferenceWithTimestamps(fairValues, usdValues, nums, dens, timestamps, count);
        } else {
            _sortMultipleByReferenceWithTimestamps(usdValues, fairValues, nums, dens, timestamps, count);
        }
    }

    function _calculateMedian(
        uint256[] memory fairValues,
        uint256[] memory usdValues,
        uint256[] memory nums,
        uint256[] memory dens,
        uint256[] memory timestamps,
        uint256 count
    ) private pure returns (MedianSet memory) {
        uint256 fairValue;
        uint256 usdValue;
        uint256 numerator;
        uint256 denominator;
        uint256 medianTimestamp;

        if (count % 2 == 1) {
            uint256 mid = count / 2;
            fairValue = fairValues[mid];
            usdValue = usdValues[mid];
            numerator = nums[mid];
            denominator = dens[mid];
            medianTimestamp = timestamps[mid];
        } else {
            uint256 mid1 = (count / 2) - 1;
            uint256 mid2 = count / 2;
            // Round to nearest by adding 1 before dividing: (a + b + 1) / 2
            // This gives proper rounding: 2.5 → 3, 2.4 → 2
            uint256 sumFair = fairValues[mid1] + fairValues[mid2];
            fairValue = (sumFair + 1) / 2;

            uint256 sumUsd = usdValues[mid1] + usdValues[mid2];
            usdValue = (sumUsd + 1) / 2;

            // Take numerator/denominator from the RIGHT middle oracle
            // This ensures consistency - all data comes from the same oracle
            numerator = nums[mid2];
            denominator = dens[mid2];

            medianTimestamp = timestamps[mid2];
        }

        return
            MedianSet({
                fairValue: fairValue,
                usdValue: usdValue,
                numerator: numerator,
                denominator: denominator,
                timestamp: medianTimestamp
            });
    }

    // Sorts main[] ascending and reorders a[], b[], c[], d[] in the same way.
    // Uses hybrid approach: optimized insertion sort for n <= 10, iterative QuickSort for n > 10
    function _sortMultipleByReferenceWithTimestamps(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 len
    ) internal pure {
        if (len <= 1) return;

        // Use optimized insertion sort for small arrays (n <= 10)
        if (len <= 10) {
            _insertionSort(main, a, b, c, d, len);
        } else {
            // Use iterative QuickSort for larger arrays (n > 10)
            _quickSortIterative(main, a, b, c, d, 0, len - 1);
        }
    }

    // Optimized insertion sort for small arrays (n <= 10)
    function _insertionSort(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 len
    ) private pure {
        for (uint256 i = 1; i < len; ++i) {
            uint256 keyMain = main[i];
            uint256 keyA = a[i];
            uint256 keyB = b[i];
            uint256 keyC = c[i];
            uint256 keyD = d[i];
            uint256 j = i;

            // Shift elements that are greater than keyMain
            while (j > 0) {
                uint256 prev = j - 1;
                if (main[prev] <= keyMain) break;

                main[j] = main[prev];
                a[j] = a[prev];
                b[j] = b[prev];
                c[j] = c[prev];
                d[j] = d[prev];
                j = prev;
            }

            main[j] = keyMain;
            a[j] = keyA;
            b[j] = keyB;
            c[j] = keyC;
            d[j] = keyD;
        }
    }

    // Iterative QuickSort implementation - eliminates recursion depth risk
    // Uses explicit stack instead of recursive calls
    // This prevents EVM stack overflow even with 1000+ oracles
    function _quickSortIterative(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure {
        // Use insertion sort for small arrays (optimization)
        if (right - left + 1 <= 10) {
            _insertionSortRange(main, a, b, c, d, left, right);
            return;
        }

        // Create an explicit stack for storing ranges to be sorted
        // Maximum stack size is O(log n), but we allocate conservatively
        // For safety with up to 1000 elements, we need at most 20 stack entries
        uint256[] memory stack = new uint256[](128); // 64 pairs of (left, right)

        uint256 stackTop = 0;

        // Push initial range
        stack[stackTop++] = left;
        stack[stackTop++] = right;

        // Process ranges until stack is empty
        while (stackTop > 0) {
            // Pop range
            right = stack[--stackTop];
            left = stack[--stackTop];

            // Use insertion sort for small partitions
            if (right - left + 1 <= 10) {
                _insertionSortRange(main, a, b, c, d, left, right);
                continue;
            }

            // Three-way partition
            (uint256 pivotEnd, uint256 pivotStart) = _partition3Way(main, a, b, c, d, left, right);

            // Push partitions onto stack for further processing
            // Partition structure:
            //   [left, pivotEnd] = elements < pivot (may be empty)
            //   [pivotEnd+1, pivotStart-1] = elements = pivot (already sorted)
            //   [pivotStart, right] = elements > pivot (may be empty)

            // Push right partition if it has 2+ elements
            // Condition: pivotStart < right means at least 2 elements exist
            if (pivotStart < right) {
                require(stackTop + 2 <= 128, "Stack overflow in iterative sort");
                stack[stackTop++] = pivotStart;
                stack[stackTop++] = right;
            }

            // Push left partition if it has 2+ elements
            // Condition: pivotEnd > left means at least 2 elements exist
            if (pivotEnd > left) {
                require(stackTop + 2 <= 128, "Stack overflow in iterative sort");
                stack[stackTop++] = left;
                stack[stackTop++] = pivotEnd;
            }
        }
    }

    // DEPRECATED: Recursive QuickSort - kept for reference but not used
    // Use _quickSortIterative instead to avoid EVM stack overflow
    function _quickSort(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure {
        if (left < right) {
            // Use insertion sort for small partitions (optimization)
            if (right - left + 1 <= 10) {
                _insertionSortRange(main, a, b, c, d, left, right);
                return;
            }

            // Three-way partition: returns (pivotEnd, pivotStart)
            // Elements in [left, pivotEnd] < pivot
            // Elements in [pivotEnd+1, pivotStart-1] = pivot
            // Elements in [pivotStart, right] > pivot
            (uint256 pivotEnd, uint256 pivotStart) = _partition3Way(main, a, b, c, d, left, right);

            // Recursively sort only the partitions that need it
            // Skip the middle partition (elements equal to pivot)
            // Check if there are elements to sort in left partition
            if (pivotEnd >= left && pivotEnd > left) {
                _quickSort(main, a, b, c, d, left, pivotEnd);
            }
            // Check if there are elements to sort in right partition
            if (pivotStart <= right && pivotStart < right) {
                _quickSort(main, a, b, c, d, pivotStart, right);
            }
        }
    }

    // Three-way partition function (Dutch National Flag algorithm)
    // Efficiently handles duplicate values by partitioning into three sections:
    // - Less than pivot
    // - Equal to pivot
    // - Greater than pivot
    // Returns (end of less-than section, start of greater-than section)
    function _partition3Way(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure returns (uint256, uint256) {
        // Use middle element as pivot
        uint256 mid = left + (right - left) / 2;
        uint256 pivot = main[mid];

        // Initialize pointers
        uint256 i = left;    // Current element
        uint256 lt = left;   // End of less-than section
        uint256 gt = right;  // Start of greater-than section

        // Move pivot to end temporarily
        _swap(main, a, b, c, d, mid, right);

        while (i <= gt) {
            if (main[i] < pivot) {
                // Element is less than pivot - swap to left section
                _swap(main, a, b, c, d, lt, i);
                lt++;
                i++;
            } else if (main[i] > pivot) {
                // Element is greater than pivot - swap to right section
                _swap(main, a, b, c, d, i, gt);
                gt--;
                // Don't increment i - need to examine the swapped element
            } else {
                // Element equals pivot - keep in middle
                i++;
            }
        }

        // lt points to first element equal to pivot
        // gt points to last element equal to pivot
        // Return:
        //   - end of less-than section (lt-1), but left if no elements less than pivot
        //   - start of greater-than section (gt+1), but right if no elements greater than pivot
        uint256 lessThanEnd = (lt > left) ? lt - 1 : left;
        uint256 greaterThanStart = (gt < right) ? gt + 1 : right;

        return (lessThanEnd, greaterThanStart);
    }

    // Swap elements at indices i and j across all arrays
    function _swap(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 i,
        uint256 j
    ) private pure {
        if (i == j) return;

        uint256 tempMain = main[i];
        uint256 tempA = a[i];
        uint256 tempB = b[i];
        uint256 tempC = c[i];
        uint256 tempD = d[i];

        main[i] = main[j];
        a[i] = a[j];
        b[i] = b[j];
        c[i] = c[j];
        d[i] = d[j];

        main[j] = tempMain;
        a[j] = tempA;
        b[j] = tempB;
        c[j] = tempC;
        d[j] = tempD;
    }

    // Insertion sort for a range (used by QuickSort for small partitions)
    function _insertionSortRange(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure {
        for (uint256 i = left + 1; i <= right; ++i) {
            uint256 keyMain = main[i];
            uint256 keyA = a[i];
            uint256 keyB = b[i];
            uint256 keyC = c[i];
            uint256 keyD = d[i];
            uint256 j = i;

            while (j > left) {
                uint256 prev = j - 1;
                if (main[prev] <= keyMain) break;

                main[j] = main[prev];
                a[j] = a[prev];
                b[j] = b[prev];
                c[j] = c[prev];
                d[j] = d[prev];
                j = prev;
            }

            main[j] = keyMain;
            a[j] = keyA;
            b[j] = keyB;
            c[j] = keyC;
            d[j] = keyD;
        }
    }

    function _parseKey(string memory key) internal pure returns (bytes32 actionHash, string memory assetKey) {
        bytes memory keyBytes = bytes(key);
        uint256 len = keyBytes.length;

        uint256 colonIndex = len;

        for (uint256 i = 0; i < len; ++i) {
            if (keyBytes[i] == bytes1(":")) {
                colonIndex = i;
                break;
            }
        }

        if (colonIndex >= len) {
            return (bytes32(0), "");
        }

        bytes memory actionBytes = new bytes(colonIndex);
        for (uint256 i = 0; i < colonIndex; ++i) {
            actionBytes[i] = keyBytes[i];
        }
        actionHash = keccak256(actionBytes);

        uint256 assetLen = len - colonIndex - 1;
        bytes memory assetBytes = new bytes(assetLen);
        for (uint256 i = 0; i < assetLen; ++i) {
            assetBytes[i] = keyBytes[colonIndex + 1 + i];
        }
        assetKey = string(assetBytes);
    }

    /// @notice Get a specific value type by parsing the key
    /// @dev Key format: "action:asset" where action is fairValue/usdValue/numerator/denominator
    /// @param key The key to parse and query
    /// @return value The requested value
    /// @return timestamp The timestamp of the value
    function getValue(string calldata key) external view returns (uint128 value, uint128 timestamp) {
        (bytes32 actionHash, string memory assetKey) = _parseKey(key);

        if (actionHash == bytes32(0)) {
            return (uint128(0), uint128(0));
        }

        MedianSet memory m = getMedianValues(assetKey);

        if (actionHash == _FAIR_VALUE) {
            return (uint128(m.fairValue), uint128(m.timestamp));
        }

        if (actionHash == _USD_VALUE) {
            return (uint128(m.usdValue), uint128(m.timestamp));
        }

        if (actionHash == _NUMERATOR) {
            return (uint128(m.numerator), uint128(m.timestamp));
        }

        if (actionHash == _DENOMINATOR) {
            return (uint128(m.denominator), uint128(m.timestamp));
        }

        return (uint128(m.fairValue), uint128(m.timestamp));
    }
}
