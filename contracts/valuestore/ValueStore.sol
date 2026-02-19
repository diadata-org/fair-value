// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";

/// @title ValueStore
/// @notice Stores fairValue, valueUsd, numerator, denominator per string key
/// @dev UUPS upgradeable contract version. Uses reinitializer for future extensibility.
contract ValueStore is Initializable, OwnableUpgradeable, UUPSUpgradeable, IERC165 {
    // --- Errors ---

    error ZeroAddress();

    // --- Value storage ---

    struct StoredValue {
        uint256 fairValue;
        uint256 valueUsd;
        uint256 numerator;
        uint256 denominator;
        uint256 timestamp;
    }

    mapping(string => StoredValue) private data;

    event ValueUpdated(
        string key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator,
        uint256 timestamp
    );

    /// @notice Storage gap for future upgrades (50 slots)
    /// @dev Storage slots from forge build --extra-output storageLayout:
    ///
    /// Slot | Label    | Type
    /// ----|----------|--------------------------------------------------
    ///   0  | data     | mapping(string => struct StoredValue)
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
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
    /// @notice Set the values for a given key (only owner)
    function setValue(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator
    ) external onlyOwner {
        require(numerator == 0 || denominator != 0, "division by zero");

        data[key] = StoredValue({
            fairValue: fairValue,
            valueUsd: valueUsd,
            numerator: numerator,
            denominator: denominator,
            timestamp: block.timestamp
        });

        emit ValueUpdated(key, fairValue, valueUsd, numerator, denominator, block.timestamp);
    }

    /// @notice Set values for multiple keys at once (only owner)
    function setMultipleValues(
        string[] calldata keys,
        uint256[] calldata fairValues,
        uint256[] calldata valueUsds,
        uint256[] calldata numerators,
        uint256[] calldata denominators
    ) external onlyOwner {
        require(
            keys.length == fairValues.length &&
                keys.length == valueUsds.length &&
                keys.length == numerators.length &&
                keys.length == denominators.length,
            "Array lengths must match"
        );

        for (uint256 i = 0; i < keys.length; i++) {
            require(numerators[i] == 0 || denominators[i] != 0, "division by zero");

            data[keys[i]] = StoredValue({
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
    function getValue(
        string calldata key
    )
        external
        view
        returns (uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
    {
        StoredValue storage sv = data[key];
        require(sv.timestamp != 0, "No data for key");
        return (sv.fairValue, sv.valueUsd, sv.numerator, sv.denominator, sv.timestamp);
    }

    /// @notice Supports ERC165 interface detection
    /// @dev Returns true for IERC165 interface and itself
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IERC165).interfaceId;
    }
}
