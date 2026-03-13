// SPDX-License-Identifier: MIT
pragma solidity 0.8.34;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";
import {IValueStore} from "../interfaces/IValueStore.sol";

/// @title ValueStore
/// @notice Stores fairValue, valueUsd, numerator, denominator per string key
/// @author DIA (diadata.org)
/// @dev UUPS upgradeable contract version. Uses reinitializer for future extensibility.
contract ValueStore is Initializable, OwnableUpgradeable, UUPSUpgradeable, IERC165, IValueStore {
    // --- Errors ---

    error ZeroAddress();
    error DivisionByZero();
    error InvalidArrayLengths();
    error NoDataForKey();
    error InvalidKey();
    error InvalidKeyInBatch(uint256 index);
    error DivisionByZeroInBatch(uint256 index);
    error TimestampTooFarInFuture(uint256 timestamp, uint256 blockTime);
    error TimestampTooFarInPast(uint256 timestamp, uint256 blockTime);
    error TimestampNotIncreasing(uint256 newTimestamp, uint256 existingTimestamp);

    // --- Value storage ---

    /// @notice Stored value data structure
    /// @dev Contains all value types for a single key along with update timestamp
    /// @param fairValue The fair value of the asset
    /// @param valueUsd The USD value of the asset
    /// @param numerator The numerator for fractional representation
    /// @param denominator The denominator for fractional representation
    /// @param timestamp The block timestamp when the value was last updated
    struct StoredValue {
        uint256 fairValue;
        uint256 valueUsd;
        uint256 numerator;
        uint256 denominator;
        uint256 timestamp;
    }

    mapping(string => StoredValue) private _data;

    /// @notice Maximum timestamp gap in the future/past (1 hour)
    uint256 public constant MAX_TIMESTAMP_GAP = 1 hours;

    /* solhint-disable gas-indexed-events */
    /// @notice Emitted when a value is updated
    /// @param key The key that was updated
    /// @param fairValue The new fair value
    /// @param valueUsd The new USD value
    /// @param numerator The new numerator
    /// @param denominator The new denominator
    /// @param timestamp The timestamp of the update
    event ValueUpdated(
        string indexed key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator,
        uint256 indexed timestamp
    );
    /* solhint-enable gas-indexed-events */

    /// @notice Storage gap for future upgrades (50 slots)
    /// @dev Storage layout (conceptual):
    ///      Parent contracts with storage: Initializable, OwnableUpgradeable
    ///      Parent contracts without storage: UUPSUpgradeable (functions only)
    ///      Interfaces: IERC165, IValueStore (no storage)
    ///
    ///      Storage order in proxy:
    ///      1. Initializable storage (_initialized, _initializing, __gap[3])
    ///      2. OwnableUpgradeable storage (_owner)
    ///      3. ValueStore storage (this contract's state variables)
    ///      4. __gap at the end (reserved for future upgrades)
    ///
    ///      IMPORTANT UPGRADE RULES:
    ///      - New state variables must be added BEFORE __gap
    ///      - Reduce __gap size by the number of new variables added
    ///      - Never remove, rename, or change the type of existing variables
    ///      - Run `forge inspect ValueStore storage-layout` for exact slot positions
    ///
    ///      Current layout (contract-relative to ValueStore):
    ///      - Slot 0: _data mapping
    ///      - Slot 1+: __gap[50] (reserve for future state variables)
    uint256[50] private __gap;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializes the contract with owner
    /// @dev Replaces constructor for upgradeable contracts. Uses reinitializer(1)
    ///      to allow future upgrades to add new initialization logic with version 2, 3, etc.
    ///      Initializes parent contracts in the correct order: Initializable, OwnableUpgradeable.
    /// @param initialOwner The address that will own the contract
    function initialize(address initialOwner) public reinitializer(1) {
        if (initialOwner == address(0)) revert ZeroAddress();
        __Ownable_init(initialOwner);
    }

    /// @notice Authorizes upgrade to new implementation
    /// @dev Only callable by contract owner
    /// @param newImplementation Address of the new implementation contract
    // solhint-disable-next-line no-empty-blocks
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /// @notice Set the values for a given key (only owner)
    /// @param key The key to set the value for
    /// @param fairValue The fair value to store
    /// @param valueUsd The USD value to store
    /// @param numerator The numerator to store
    /// @param denominator The denominator to store
    /// @param timestamp The timestamp to store
    /// @dev Emits ValueUpdated event with provided timestamp
    function setValue(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator,
        uint256 timestamp
    ) external onlyOwner {
        if (bytes(key).length == 0) {
            revert InvalidKey();
        }
        if (numerator != 0 && denominator == 0) {
            revert DivisionByZero();
        }

        _validateTimestamp(key, timestamp);

        _data[key] = StoredValue({
            fairValue: fairValue,
            valueUsd: valueUsd,
            numerator: numerator,
            denominator: denominator,
            timestamp: timestamp
        });

        emit ValueUpdated(key, fairValue, valueUsd, numerator, denominator, timestamp);
    }

    /// @notice Set values for multiple keys at once (only owner)
    /// @param keys Array of keys to set values for
    /// @param fairValues Array of fair values to store
    /// @param valueUsds Array of USD values to store
    /// @param numerators Array of numerators to store
    /// @param denominators Array of denominators to store
    /// @param timestamps Array of timestamps to store
    /// @dev All arrays must have the same length. Emits ValueUpdated event for each key.
    function setMultipleValues(
        string[] calldata keys,
        uint256[] calldata fairValues,
        uint256[] calldata valueUsds,
        uint256[] calldata numerators,
        uint256[] calldata denominators,
        uint256[] calldata timestamps
    ) external onlyOwner {
        uint256 length = keys.length;
        if (
            length != fairValues.length ||
            length != valueUsds.length ||
            length != numerators.length ||
            length != denominators.length ||
            length != timestamps.length
        ) {
            revert InvalidArrayLengths();
        }

        for (uint256 i = 0; i < length; ++i) {
            if (bytes(keys[i]).length == 0) {
                revert InvalidKeyInBatch(i);
            }
            if (numerators[i] != 0 && denominators[i] == 0) {
                revert DivisionByZeroInBatch(i);
            }

            _validateTimestamp(keys[i], timestamps[i]);

            _data[keys[i]] = StoredValue({
                fairValue: fairValues[i],
                valueUsd: valueUsds[i],
                numerator: numerators[i],
                denominator: denominators[i],
                timestamp: timestamps[i]
            });

            emit ValueUpdated(keys[i], fairValues[i], valueUsds[i], numerators[i], denominators[i], timestamps[i]);
        }
    }

    /// @notice Get the stored values for a given key
    /// @param key The key to retrieve values for
    /// @return fairValue The stored fair value
    /// @return valueUsd The stored USD value
    /// @return numerator The stored numerator
    /// @return denominator The stored denominator
    /// @return timestamp The timestamp of the last update
    function getValue(
        string calldata key
    )
        external
        view
        returns (uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
    {
        StoredValue storage sv = _data[key];
        if (sv.timestamp == 0) {
            revert NoDataForKey();
        }
        return (sv.fairValue, sv.valueUsd, sv.numerator, sv.denominator, sv.timestamp);
    }

    /// @notice Supports ERC165 interface detection
    /// @dev Returns true for IERC165 and IValueStore interfaces
    /// @param interfaceId The interface identifier to check
    /// @return bool True if the contract supports the interface
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            interfaceId == type(IERC165).interfaceId ||
            interfaceId == type(IValueStore).interfaceId;
    }

    /// @notice Validates that a timestamp is within acceptable bounds.
    /// @dev Timestamp must not be too far in the future or too far in the past.
    ///      This prevents invalid data from polluting the oracle.
    ///      Also ensures timestamps are monotonically increasing for each key.
    /// @param key The asset identifier to check existing timestamp for.
    /// @param timestamp The timestamp to validate.
    function _validateTimestamp(string memory key, uint256 timestamp) private view {
        uint256 currentBlockTime = block.timestamp;

        // Check if timestamp is too far in the future
        if (timestamp > currentBlockTime + MAX_TIMESTAMP_GAP) {
            revert TimestampTooFarInFuture(timestamp, currentBlockTime);
        }

        // Check if timestamp is too far in the past
        if (currentBlockTime > MAX_TIMESTAMP_GAP && timestamp < currentBlockTime - MAX_TIMESTAMP_GAP) {
            revert TimestampTooFarInPast(timestamp, currentBlockTime);
        }

        // Ensure timestamp is not older than existing value for this key
        StoredValue storage existingValue = _data[key];
        if (existingValue.timestamp != 0) {
            if (timestamp <= existingValue.timestamp) {
                revert TimestampNotIncreasing(timestamp, existingValue.timestamp);
            }
        }
    }
}
