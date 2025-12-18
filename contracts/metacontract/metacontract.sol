// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.29;

interface IValueStore {
    function getValue(
        string memory key
    )
        external
        view
        returns (
            uint256 fairValue,
            uint256 usdValue,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        );
}

library QuickSort {
    function sort(
        uint256[] memory arr,
        uint256 left,
        uint256 right
    ) internal pure {
        if (left >= right) return;
        uint256 i = left;
        uint256 j = right;
        uint256 pivot = arr[left + (right - left) / 2];
        while (i <= j) {
            while (arr[i] < pivot) i++;
            while (arr[j] > pivot) j--;
            if (i <= j) {
                (arr[i], arr[j]) = (arr[j], arr[i]);
                i++;
                if (j > 0) j--;
            }
        }
        if (left < j) sort(arr, left, j);
        if (i < right) sort(arr, i, right);
    }
}

contract ValueStoreMetaAggregator {
    using QuickSort for uint256[];

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
    address public owner;

    bytes32 private constant FAIR_VALUE = keccak256("fairValue");
    bytes32 private constant USD_VALUE = keccak256("usdValue");
    bytes32 private constant NUMERATOR = keccak256("numerator");
    bytes32 private constant DENOMINATOR = keccak256("denominator");

    error InvalidThreshold(uint256);
    error InvalidTimeout(uint256);
    error ThresholdNotMet(uint256, uint256);
    error ZeroAddress();
    error OracleExists();
    error OracleNotFound();

    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    constructor(address _owner) {
        owner = _owner;
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
        threshold = newThreshold;
    }

    function setTimeoutSeconds(uint256 newTimeout) external onlyOwner {
        if (newTimeout == 0) revert InvalidTimeout(newTimeout);
        timeoutSeconds = newTimeout;
    }

    function getMedianValues(
        string memory key
    ) external view returns (MedianSet memory) {
        if (threshold == 0) revert InvalidThreshold(threshold);
        if (timeoutSeconds == 0) revert InvalidTimeout(timeoutSeconds);

        uint256[] memory fairValues = new uint256[](numValueStores);
        uint256[] memory usdValues = new uint256[](numValueStores);
        uint256[] memory nums = new uint256[](numValueStores);
        uint256[] memory dens = new uint256[](numValueStores);
        uint256 count = 0;

        for (uint256 i = 0; i < numValueStores; ++i) {
            IValueStore store = IValueStore(valueStores[i]);
            (
                uint256 fairV,
                uint256 usdV,
                uint256 num,
                uint256 den,
                uint256 ts
            ) = store.getValue(key);

            if (ts + timeoutSeconds < block.timestamp) continue;

            fairValues[count] = fairV;
            usdValues[count] = usdV;
            nums[count] = num;
            dens[count] = den;
            count++;
        }

        if (count < threshold) revert ThresholdNotMet(count, threshold);

        uint256 last = count - 1;
        fairValues.sort(0, last);
        usdValues.sort(0, last);
        nums.sort(0, last);
        dens.sort(0, last);

        uint256 mid = count / 2;

        return
            MedianSet({
                fairValue: fairValues[mid],
                usdValue: usdValues[mid],
                numerator: nums[mid],
                denominator: dens[mid],
                timestamp: block.timestamp
            });
    }

    function _parseKey(
        string memory key
    ) internal pure returns (bytes32 actionHash, string memory assetKey) {
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

    function getValue(
        string memory key
    ) external view returns (uint128, uint128) {
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
