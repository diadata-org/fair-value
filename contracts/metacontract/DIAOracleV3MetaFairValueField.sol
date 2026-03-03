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

    bytes32 private constant _FAIR_VALUE = keccak256("fairValue");
    bytes32 private constant _USD_VALUE = keccak256("usdValue");
    bytes32 private constant _NUMERATOR = keccak256("numerator");
    bytes32 private constant _DENOMINATOR = keccak256("denominator");
    uint256 private constant _MAX_TIMEOUT_SECONDS = 1 days;

    error InvalidThreshold(uint256);
    error InvalidTimeOut(uint256);
    error TimeoutExceedsLimit(uint256);
    error ThresholdNotMet(uint256, uint256);
    error ZeroAddress();
    error OracleExists();
    error OracleNotFound();

    /// @notice Emitted when timeout seconds is changed
    /// @param oldTimeoutSeconds The previous timeout value
    /// @param newTimeoutSeconds The new timeout value
    event TimeoutSecondsChanged(uint256 indexed oldTimeoutSeconds, uint256 indexed newTimeoutSeconds);

    /// @notice Emitted when threshold is changed
    /// @param oldThreshold The previous threshold value
    /// @param newThreshold The new threshold value
    event ThresholdChanged(uint256 indexed oldThreshold, uint256 indexed newThreshold);

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
    }

    /// @notice Add a new ValueStore to the oracle
    /// @param newStore The address of the ValueStore to add
    function addValueStore(address newStore) external onlyOwner {
        if (newStore == address(0)) revert ZeroAddress();
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
            fairValue = (fairValues[mid1] + fairValues[mid2]) / 2;
            usdValue = (usdValues[mid1] + usdValues[mid2]) / 2;
            numerator = (nums[mid1] + nums[mid2]) / 2;
            denominator = (dens[mid1] + dens[mid2]) / 2;
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
    // Uses hybrid approach: optimized insertion sort for n <= 10, QuickSort for n > 10
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
            // Use QuickSort for larger arrays (n > 10)
            _quickSort(main, a, b, c, d, 0, len - 1);
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

    // QuickSort implementation for larger arrays (n > 10)
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

            uint256 pivotIndex = _partition(main, a, b, c, d, left, right);

            // Recursively sort elements before and after partition
            if (pivotIndex > 0) {
                _quickSort(main, a, b, c, d, left, pivotIndex - 1);
            }
            _quickSort(main, a, b, c, d, pivotIndex + 1, right);
        }
    }

    // Partition function for QuickSort
    function _partition(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256[] memory d,
        uint256 left,
        uint256 right
    ) private pure returns (uint256) {
        // Use middle element as pivot (better than left/right for sorted data)
        uint256 mid = left + (right - left) / 2;
        uint256 pivot = main[mid];

        // Move pivot to end
        _swap(main, a, b, c, d, mid, right);

        uint256 storeIndex = left;
        for (uint256 i = left; i < right; ++i) {
            if (main[i] < pivot) {
                _swap(main, a, b, c, d, i, storeIndex);
                ++storeIndex;
            }
        }

        // Move pivot to its final place
        _swap(main, a, b, c, d, storeIndex, right);
        return storeIndex;
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
