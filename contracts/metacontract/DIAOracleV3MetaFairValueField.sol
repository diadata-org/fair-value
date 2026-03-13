// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.34;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";
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
    error InvalidMaxValueStores(uint256);
    error InvalidTimeOut(uint256);
    error TimeoutExceedsLimit(uint256);
    error ThresholdNotMet(uint256, uint256);
    error ZeroAddress();
    error OracleExists();
    error OracleNotFound();
    error MaxValueStoresReached(uint256);
    error InvalidValueStoreInterface(address store);
    error StackOverflow();
    error UnrecognizedAction();

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
    /// @dev Reverts if the store doesn't implement IValueStore interface
    function addValueStore(address newStore) external onlyOwner {
        if (newStore == address(0)) revert ZeroAddress();
        if (numValueStores == maxValueStores) revert MaxValueStoresReached(maxValueStores);
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == newStore) revert OracleExists();
        }

        // Verify that the store implements IValueStore interface via ERC165
        if (!_supportsIValueStore(newStore)) {
            revert InvalidValueStoreInterface(newStore);
        }

        ++numValueStores;
        valueStores[numValueStores - 1] = newStore;
        emit ValueStoreAdded(newStore, numValueStores - 1);
    }

    /// @notice Remove a ValueStore from the oracle
    /// @param storeAddr The address of the ValueStore to remove
    /// @dev Reverts if the store is not found
    function removeValueStore(address storeAddr) external onlyOwner {
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == storeAddr) {
                --numValueStores;
                valueStores[i] = valueStores[numValueStores];
                emit ValueStoreRemoved(storeAddr, i);
                delete valueStores[numValueStores];
                return;
            }
        }
        revert OracleNotFound();
    }

    /// @notice Check if a contract supports the IValueStore interface
    /// @param store The address to check
    /// @return bool True if the contract supports IValueStore interface
    function _supportsIValueStore(address store) private view returns (bool) {
        // First check if address has code (EOAs will fail this check)
        if (store.code.length == 0) {
            return false;
        }

        // Check if address supports ERC165 and IValueStore interface
        try IERC165(store).supportsInterface(type(IValueStore).interfaceId) returns (bool supported) {
            return supported;
        } catch {
            return false;
        }
    }

    /// @notice Set the minimum number of valid responses required
    /// @param newThreshold The new threshold value
    /// @dev Reverts if newThreshold is zero
    function setThreshold(uint256 newThreshold) external onlyOwner {
        if (newThreshold == 0) revert InvalidThreshold(newThreshold);
        uint256 oldThreshold = threshold;
        threshold = newThreshold;
        emit ThresholdChanged(oldThreshold, newThreshold);
    }

    /// @notice Set the maximum age of data considered valid
    /// @param newTimeoutSeconds The new timeout in seconds
    /// @dev Reverts if newTimeoutSeconds is zero or exceeds _MAX_TIMEOUT_SECONDS
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
    /// @dev Reverts if newMaxValueStores is zero or less than current count
    function setMaxValueStores(uint256 newMaxValueStores) external onlyOwner {
        if (newMaxValueStores == 0) revert InvalidMaxValueStores(newMaxValueStores);
        if (newMaxValueStores < numValueStores) {
            revert InvalidMaxValueStores(newMaxValueStores); // Cannot set below current count
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

        bool isFairValueSort = _sortValues(fairValues, usdValues, nums, dens, timestamps, count);

        return _calculateMedian(fairValues, usdValues, nums, dens, timestamps, count, isFairValueSort);
    }

    /// @notice Initialize value arrays for collecting data from stores
    /// @return fairValues Array to store fair values
    /// @return usdValues Array to store USD values
    /// @return nums Array to store numerators
    /// @return dens Array to store denominators
    /// @return timestamps Array to store timestamps
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

    /// @notice Collect values from all stores, filtering out stale data and failures
    /// @param key The key to query from stores
    /// @param fairValues Array to fill with fair values
    /// @param usdValues Array to fill with USD values
    /// @param nums Array to fill with numerators
    /// @param dens Array to fill with denominators
    /// @param timestamps Array to fill with timestamps
    /// @return count The number of valid responses collected
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
                ++count;
            } catch {
                continue;
            }
        }
    }

    /// @notice Ensure the minimum number of valid responses was collected
    /// @param count The number of valid responses collected
    /// @dev Reverts if count is less than threshold
    function _ensureThresholdMet(uint256 count) private view {
        if (count < threshold) revert ThresholdNotMet(count, threshold);
    }

    /// @notice Sort all value arrays by fairValue using iterative quicksort
    /// @param fairValues Array of fair values to sort by
    /// @param usdValues Array of USD values (sorted in parallel)
    /// @param nums Array of numerators (sorted in parallel)
    /// @param dens Array of denominators (sorted in parallel)
    /// @param timestamps Array of timestamps (sorted in parallel)
    /// @param count Number of elements to sort
    function _sortValues(
        uint256[] memory fairValues,
        uint256[] memory usdValues,
        uint256[] memory nums,
        uint256[] memory dens,
        uint256[] memory timestamps,
        uint256 count
    ) private pure returns (bool isFairValueSort) {
        if (count == 0) return true; // Default/fallback: fairValue sort

        uint256 fairSum = 0;
        for (uint256 i = 0; i < count; ++i) {
            fairSum += fairValues[i];
        }

        if (fairSum != 0) {
            _sortMultipleByReferenceWithTimestamps(fairValues, usdValues, nums, dens, timestamps, count);
            return true; // Sorted by fairValues
        } else {
            _sortMultipleByReferenceWithTimestamps(usdValues, fairValues, nums, dens, timestamps, count);
            return false; // Sorted by usdValues
        }
    }

   
    error AllPrincipalEntriesZero();

    /// @notice Calculate the median of sorted value arrays
    /// @param fairValues Sorted array of fair values
    /// @param usdValues Sorted array of USD values
    /// @param nums Sorted array of numerators
    /// @param dens Sorted array of denominators
    /// @param timestamps Sorted array of timestamps
    /// @param count Number of elements to consider
    /// @return median MedianSet containing median values and timestamp
    function _calculateMedian(
        uint256[] memory fairValues,
        uint256[] memory usdValues,
        uint256[] memory nums,
        uint256[] memory dens,
        uint256[] memory timestamps,
        uint256 count,
        bool isFairValueSort
    ) private pure returns (MedianSet memory) {
        uint256[] memory principalArray = isFairValueSort ? fairValues : usdValues;
        uint256 startIdx = 0;

        // Skip leading zeros in the principalArray after sorting
        while (startIdx < count && principalArray[startIdx] == 0) {
            startIdx++;
        }

        uint256 filteredCount = count - startIdx;
        if (filteredCount == 0) {
            revert AllPrincipalEntriesZero();
        }

        uint256 mid1;
        uint256 mid2;
        uint256 arrayMid = filteredCount / 2;
        uint256 base = startIdx;

        if (filteredCount % 2 == 1) {
            uint256 mid = base + arrayMid;
            return MedianSet(
                fairValues[mid],
                usdValues[mid],
                nums[mid],
                dens[mid],
                timestamps[mid]
            );
        } else {
            mid1 = base + arrayMid - 1;
            mid2 = base + arrayMid;
            return MedianSet(
                (fairValues[mid1] + fairValues[mid2] + 1) / 2,
                (usdValues[mid1] + usdValues[mid2] + 1) / 2,
                (nums[mid1] + nums[mid2] + 1) / 2,
                (dens[mid1] + dens[mid2] + 1) / 2,
                timestamps[mid2] // Use second value's timestamp
            );
        }
    }

    /// @notice Sort main array ascending and reorder auxiliary arrays in parallel
    /// @dev Uses hybrid approach: insertion sort for n <= 10, iterative QuickSort for n > 10
    /// @param main Primary array to sort by (ascending order)
    /// @param a Auxiliary array 1 (reordered in parallel)
    /// @param b Auxiliary array 2 (reordered in parallel)
    /// @param c Auxiliary array 3 (reordered in parallel)
    /// @param d Auxiliary array 4 (reordered in parallel)
    /// @param len Number of elements to sort
    function _sortMultipleByReferenceWithTimestamps(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 len
    ) internal pure {
        if (len < 2) return;

        // Use optimized insertion sort for small arrays (n <= 10)
        if (len < 11) {
            _insertionSort(main, a, b, c, d, len);
        } else {
            // Use iterative QuickSort for larger arrays (n > 10)
            _quickSortIterative(main, a, b, c, d, 0, len - 1);
        }
    }

    /// @notice Optimized insertion sort for small arrays (n <= 10)
    /// @param main Primary array to sort
    /// @param a Auxiliary array 1 (reordered in parallel)
    /// @param b Auxiliary array 2 (reordered in parallel)
    /// @param c Auxiliary array 3 (reordered in parallel)
    /// @param d Auxiliary array 4 (reordered in parallel)
    /// @param len Number of elements to sort
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
                if (main[prev] < keyMain + 1) break;

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

    /// @notice Iterative QuickSort implementation - eliminates recursion depth risk
    /// @dev Uses explicit stack instead of recursive calls to prevent EVM stack overflow
    /// @param main Primary array to sort
    /// @param a Auxiliary array 1 (reordered in parallel)
    /// @param b Auxiliary array 2 (reordered in parallel)
    /// @param c Auxiliary array 3 (reordered in parallel)
    /// @param d Auxiliary array 4 (reordered in parallel)
    /// @param left Left boundary of range to sort
    /// @param right Right boundary of range to sort
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
        if (right - left < 10) {
            _insertionSortRange(main, a, b, c, d, left, right);
            return;
        }

        // Create an explicit stack for storing ranges to be sorted
        // Maximum stack size is O(log n), but we allocate conservatively
        // For safety with up to 1000 elements, we need at most 20 stack entries
        uint256[] memory stack = new uint256[](128); // 64 pairs of (left, right)

        uint256 stackTop = 0;

        // Push initial range
        stack[stackTop] = left;
        ++stackTop;
        stack[stackTop] = right;
        ++stackTop;

        // Process ranges until stack is empty
        while (stackTop > 0) {
            // Pop range
            --stackTop;
            right = stack[stackTop];
            --stackTop;
            left = stack[stackTop];

            // Use insertion sort for small partitions
            if (right - left < 10) {
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
                if (stackTop + 1 > 127) revert StackOverflow();
                stack[stackTop] = pivotStart;
                ++stackTop;
                stack[stackTop] = right;
                ++stackTop;
            }

            // Push left partition if it has 2+ elements
            // Condition: pivotEnd > left means at least 2 elements exist
            if (pivotEnd > left) {
                if (stackTop + 1 > 127) revert StackOverflow();
                stack[stackTop] = left;
                ++stackTop;
                stack[stackTop] = pivotEnd;
                ++stackTop;
            }
        }
    }

    /// @notice Three-way partition (Dutch National Flag algorithm)
    /// @dev Efficiently handles duplicate values by partitioning into <, =, > sections
    /// @param main Primary array to partition
    /// @param a Auxiliary array 1 (reordered in parallel)
    /// @param b Auxiliary array 2 (reordered in parallel)
    /// @param c Auxiliary array 3 (reordered in parallel)
    /// @param d Auxiliary array 4 (reordered in parallel)
    /// @param left Left boundary of range to partition
    /// @param right Right boundary of range to partition
    /// @return lessThanEnd End index of less-than section
    /// @return greaterThanStart Start index of greater-than section
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

        while (i < gt + 1) {
            if (main[i] < pivot) {
                // Element is less than pivot - swap to left section
                _swap(main, a, b, c, d, lt, i);
                ++lt;
                ++i;
            } else if (main[i] > pivot) {
                // Element is greater than pivot - swap to right section
                _swap(main, a, b, c, d, i, gt);
                --gt;
                // Don't increment i - need to examine the swapped element
            } else {
                // Element equals pivot - keep in middle
                ++i;
            }
        }

        // lt points to first element equal to pivot
        // gt points to last element equal to pivot
        // Return:
        //   - end of less-than section (lt-1), but left if no elements less than pivot
        //   - start of greater-than section (gt+1), but right if no elements greater than pivot
        uint256 lessThanEnd = (left < lt) ? lt - 1 : left;
        uint256 greaterThanStart = (gt < right) ? gt + 1 : right;

        return (lessThanEnd, greaterThanStart);
    }

    /// @notice Swap elements at indices i and j across all arrays
    /// @param main Primary array
    /// @param a Auxiliary array 1
    /// @param b Auxiliary array 2
    /// @param c Auxiliary array 3
    /// @param d Auxiliary array 4
    /// @param i First index to swap
    /// @param j Second index to swap
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

    /// @notice Insertion sort for a range (used by QuickSort for small partitions)
    /// @param main Primary array to sort
    /// @param a Auxiliary array 1 (reordered in parallel)
    /// @param b Auxiliary array 2 (reordered in parallel)
    /// @param c Auxiliary array 3 (reordered in parallel)
    /// @param d Auxiliary array 4 (reordered in parallel)
    /// @param left Left boundary of range to sort
    /// @param right Right boundary of range to sort
    function _insertionSortRange(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure {
        for (uint256 i = left + 1; i < right + 1; ++i) {
            uint256 keyMain = main[i];
            uint256 keyA = a[i];
            uint256 keyB = b[i];
            uint256 keyC = c[i];
            uint256 keyD = d[i];
            uint256 j = i;

            while (j > left) {
                uint256 prev = j - 1;
                if (main[prev] < keyMain + 1) break;

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

    /// @notice Parse key into action hash and asset key
    /// @dev Expected format: "action:asset" (e.g., "fairValue:BTC/USD")
    /// @param key The key to parse
    /// @return actionHash Keccak256 hash of the action (fairValue/usdValue/numerator/denominator)
    /// @return assetKey The asset portion of the key (after the colon)
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

        if (colonIndex == len) {
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
            revert UnrecognizedAction();
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

        revert UnrecognizedAction();
    }

    /// @notice Get all registered ValueStore addresses
    /// @dev Returns an array of all currently registered ValueStore addresses
    /// @return stores Array of ValueStore addresses
    function getValueStores() external view returns (address[] memory stores) {
        stores = new address[](numValueStores);
        for (uint256 i = 0; i < numValueStores; ++i) {
            stores[i] = valueStores[i];
        }
    }
}
