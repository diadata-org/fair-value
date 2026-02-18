// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.30;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

interface IValueStore {
    function getValue(
        string memory key
    )
        external
        view
        returns (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp);
}

contract DIAOracleV3MetaFairValueField is Ownable {
    struct MedianSet {
        uint256 fairValue;
        uint256 usdValue;
        uint256 numerator;
        uint256 denominator;
        uint256 timestamp;
    }

    mapping(uint256 => address) public valueStores;
    uint256 public numValueStores;
    uint256 public threshold;
    uint256 public timeoutSeconds;

    bytes32 private constant FAIR_VALUE = keccak256("fairValue");
    bytes32 private constant USD_VALUE = keccak256("usdValue");
    bytes32 private constant NUMERATOR = keccak256("numerator");
    bytes32 private constant DENOMINATOR = keccak256("denominator");

    error InvalidThreshold(uint256);
    error InvalidTimeOut(uint256);
    error TimeoutExceedsLimit(uint256);
    error ThresholdNotMet(uint256, uint256);
    error ZeroAddress();
    error OracleExists();
    error OracleNotFound();

    event TimeoutSecondsChanged(uint256 indexed oldTimeoutSeconds, uint256 indexed newTimeoutSeconds);
    event ThresholdChanged(uint256 indexed oldThreshold, uint256 indexed newThreshold);

    constructor(address initialOwner) Ownable(initialOwner) {
        if (initialOwner == address(0)) revert ZeroAddress();
    }

    function addValueStore(address newStore) external onlyOwner {
        if (newStore == address(0)) revert ZeroAddress();
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == newStore) revert OracleExists();
        }
        valueStores[numValueStores++] = newStore;
    }

    function removeValueStore(address storeAddr) external onlyOwner {
        for (uint256 i = 0; i < numValueStores; ++i) {
            if (valueStores[i] == storeAddr) {
                valueStores[i] = valueStores[numValueStores - 1];
                delete valueStores[numValueStores - 1];
                numValueStores--;
                return;
            }
        }
        revert OracleNotFound();
    }

    function setThreshold(uint256 newThreshold) external onlyOwner {
        if (newThreshold == 0) revert InvalidThreshold(newThreshold);
        uint256 oldThreshold = threshold;
        threshold = newThreshold;
        emit ThresholdChanged(oldThreshold, newThreshold);
    }

    function setTimeoutSeconds(uint256 newTimeoutSeconds) external onlyOwner {
        if (newTimeoutSeconds == 0) {
            revert InvalidTimeOut(newTimeoutSeconds);
        }
        if (newTimeoutSeconds > 86400) {
            revert TimeoutExceedsLimit(newTimeoutSeconds);
        }
        uint256 oldTimeoutSeconds = timeoutSeconds;
        timeoutSeconds = newTimeoutSeconds;
        emit TimeoutSecondsChanged(oldTimeoutSeconds, newTimeoutSeconds);
    }

    // --- CORRECTED MEDIAN LOGIC BELOW ---
    function getMedianValues(string memory key) external view returns (MedianSet memory) {
        if (threshold == 0) revert InvalidThreshold(threshold);
        if (timeoutSeconds == 0) revert InvalidTimeOut(timeoutSeconds);

        uint256[] memory fairValues = new uint256[](numValueStores);
        uint256[] memory usdValues = new uint256[](numValueStores);
        uint256[] memory nums = new uint256[](numValueStores);
        uint256[] memory dens = new uint256[](numValueStores);
        uint256 count = 0;

        for (uint256 i = 0; i < numValueStores; ++i) {
            IValueStore store = IValueStore(valueStores[i]);
            try store.getValue(key) returns (uint256 fairV, uint256 usdV, uint256 num, uint256 den, uint256 ts) {
                if (ts + timeoutSeconds < block.timestamp) continue;

                fairValues[count] = fairV;
                usdValues[count] = usdV;
                nums[count] = num;
                dens[count] = den;
                count++;
            } catch {
                continue;
            }
        }

        if (count < threshold) revert ThresholdNotMet(count, threshold);

        // Calculate sum of fairValues
        uint256 fairSum = 0;
        for (uint256 i = 0; i < count; i++) {
            fairSum += fairValues[i];
        }

        // Sort: preserve correspondence across all value arrays
        if (count > 0) {
            if (fairSum != 0) {
                sortMultipleByReference(fairValues, usdValues, nums, dens, count);
            } else {
                sortMultipleByReference(usdValues, fairValues, nums, dens, count);
            }
        }

        // Calculate true median
        uint256 fairValue;
        uint256 usdValue;
        uint256 numerator;
        uint256 denominator;

        if (count % 2 == 1) {
            uint256 mid = count / 2;
            fairValue = fairValues[mid];
            usdValue = usdValues[mid];
            numerator = nums[mid];
            denominator = dens[mid];
        } else {
            uint256 mid1 = (count / 2) - 1;
            uint256 mid2 = count / 2;
            fairValue = (fairValues[mid1] + fairValues[mid2]) / 2;
            usdValue = (usdValues[mid1] + usdValues[mid2]) / 2;
            numerator = (nums[mid1] + nums[mid2]) / 2;
            denominator = (dens[mid1] + dens[mid2]) / 2;
        }

        return MedianSet({
            fairValue: fairValue,
            usdValue: usdValue,
            numerator: numerator,
            denominator: denominator,
            timestamp: block.timestamp
        });
    }

    // Sorts main[] ascending and reorders a[], b[], c[] in the same way.
    function sortMultipleByReference(
        uint256[] memory main,
        uint256[] memory a,
        uint256[] memory b,
        uint256[] memory c,
        uint256 len
    ) internal pure {
        for (uint256 i = 1; i < len; i++) {
            uint256 keyMain = main[i];
            uint256 keyA = a[i];
            uint256 keyB = b[i];
            uint256 keyC = c[i];
            uint256 j = i;
            while (j != 0 && main[j - 1] > keyMain) {
                main[j] = main[j - 1];
                a[j] = a[j - 1];
                b[j] = b[j - 1];
                c[j] = c[j - 1];
                j--;
            }
            main[j] = keyMain;
            a[j] = keyA;
            b[j] = keyB;
            c[j] = keyC;
        }
    }

    function _parseKey(string memory key) internal pure returns (bytes32 actionHash, string memory assetKey) {
        bytes memory keyBytes = bytes(key);
        uint256 len = keyBytes.length;

        uint256 colonIndex = len;

        for (uint256 i = 0; i < len; i++) {
            if (keyBytes[i] == bytes1(":")) {
                colonIndex = i;
                break;
            }
        }

        if (colonIndex >= len) {
            return (bytes32(0), "");
        }

        bytes memory actionBytes = new bytes(colonIndex);
        for (uint256 i = 0; i < colonIndex; i++) {
            actionBytes[i] = keyBytes[i];
        }
        actionHash = keccak256(actionBytes);

        uint256 assetLen = len - colonIndex - 1;
        bytes memory assetBytes = new bytes(assetLen);
        for (uint256 i = 0; i < assetLen; i++) {
            assetBytes[i] = keyBytes[colonIndex + 1 + i];
        }
        assetKey = string(assetBytes);
    }

    function getValue(string memory key) external view returns (uint128, uint128) {
        (bytes32 actionHash, string memory assetKey) = _parseKey(key);

        if (actionHash == bytes32(0)) {
            return (uint128(0), uint128(0));
        }

        MedianSet memory m = this.getMedianValues(assetKey);

        if (actionHash == FAIR_VALUE) {
            return (uint128(m.fairValue), uint128(m.timestamp));
        }

        if (actionHash == USD_VALUE) {
            return (uint128(m.usdValue), uint128(m.timestamp));
        }

        if (actionHash == NUMERATOR) {
            return (uint128(m.numerator), uint128(m.timestamp));
        }

        if (actionHash == DENOMINATOR) {
            return (uint128(m.denominator), uint128(m.timestamp));
        }

        return (uint128(m.fairValue), uint128(m.timestamp));
    }
}