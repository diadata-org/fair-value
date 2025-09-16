// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title ValueStore
/// @notice Stores fairValue, valueUsd, numerator, denominator per string key, only owner can update
contract ValueStore {

    // --- Ownable pattern (like in the DIA contract) ---
    address public owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    modifier onlyOwner() {
        require(msg.sender == owner, "Caller is not the owner");
        _;
    }

    constructor() {
        owner = msg.sender;
        emit OwnershipTransferred(address(0), msg.sender);
    }

    function transferOwnership(address newOwner) external onlyOwner {
        require(newOwner != address(0), "New owner is zero address");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

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

    //  TO DO: Should we add timestamp as input var?
    /// @notice Set the values for a given key (only owner)
    function setValue(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator
    ) external onlyOwner {
        require(denominator != 0, "Denominator cannot be zero");

        data[key] = StoredValue({
            fairValue: fairValue,
            valueUsd: valueUsd,
            numerator: numerator,
            denominator: denominator,
            timestamp: block.timestamp
        });

        emit ValueUpdated(
            key,
            fairValue,
            valueUsd,
            numerator,
            denominator,
            block.timestamp
        );
    }

    //  TO DO: Should we add timestamps as input var?
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
            require(denominators[i] != 0, "Denominator cannot be zero");

            data[keys[i]] = StoredValue({
                fairValue: fairValues[i],
                valueUsd: valueUsds[i],
                numerator: numerators[i],
                denominator: denominators[i],
                timestamp: block.timestamp
            });

            emit ValueUpdated(
                keys[i],
                fairValues[i],
                valueUsds[i],
                numerators[i],
                denominators[i],
                block.timestamp
            );
        }
    }

    /// @notice Get the stored values for a given key
    function getValue(string calldata key)
        external
        view
        returns (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        )
    {
        StoredValue storage sv = data[key];
        require(sv.timestamp != 0, "No data for key");
        return (
            sv.fairValue,
            sv.valueUsd,
            sv.numerator,
            sv.denominator,
            sv.timestamp
        );
    }
}
