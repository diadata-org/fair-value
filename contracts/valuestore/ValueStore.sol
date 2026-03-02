// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

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

    // --- Value storage ---

    struct StoredValue {
        uint256 fairValue;
        uint256 valueUsd;
        uint256 numerator;
        uint256 denominator;
        uint256 timestamp;
    }

    mapping(string => StoredValue) private _data;

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
    /// @dev Storage slots from forge build --extra-output storageLayout:
    ///
    /// Slot | Label    | Type
    /// ----|----------|--------------------------------------------------
    ///   0  | _data    | mapping(string => struct StoredValue)
    ///   1+ | __gap   | 50 slots reserved for future upgrades
    ///
    /// Note: Parent contract storage (before slot 0):
    /// - Initializable: 5 slots (_initialized, _initializing, __gap)
    /// - OwnableUpgradeable: 1 slot (_owner)
    /// Total contract uses slot 0 + parent slots + 50 gap slots
    uint256[50] private __gap;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializes the contract with owner
    /// @dev Replaces constructor for upgradeable contracts. Uses reinitializer(1)
    ///      to allow future upgrades to add new initialization logic with version 2, 3, etc.
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
    function setValue(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator
    ) external onlyOwner {
        if (numerator != 0 && denominator == 0) {
            revert DivisionByZero();
        }

        _data[key] = StoredValue({
            fairValue: fairValue,
            valueUsd: valueUsd,
            numerator: numerator,
            denominator: denominator,
            timestamp: block.timestamp
        });

        emit ValueUpdated(key, fairValue, valueUsd, numerator, denominator, block.timestamp);
    }

    /// @notice Set values for multiple keys at once (only owner)
    /// @param keys Array of keys to set values for
    /// @param fairValues Array of fair values to store
    /// @param valueUsds Array of USD values to store
    /// @param numerators Array of numerators to store
    /// @param denominators Array of denominators to store
    function setMultipleValues(
        string[] calldata keys,
        uint256[] calldata fairValues,
        uint256[] calldata valueUsds,
        uint256[] calldata numerators,
        uint256[] calldata denominators
    ) external onlyOwner {
        uint256 length = keys.length;
        if (
            length != fairValues.length ||
            length != valueUsds.length ||
            length != numerators.length ||
            length != denominators.length
        ) {
            revert InvalidArrayLengths();
        }

        for (uint256 i = 0; i < length; ++i) {
            if (numerators[i] != 0 && denominators[i] == 0) {
                revert DivisionByZero();
            }

            _data[keys[i]] = StoredValue({
                fairValue: fairValues[i],
                valueUsd: valueUsds[i],
                numerator: numerators[i],
                denominator: denominators[i],
                timestamp: block.timestamp
            });

            emit ValueUpdated(keys[i], fairValues[i], valueUsds[i], numerators[i], denominators[i], block.timestamp);
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
}
