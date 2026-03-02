// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/// @title IValueStore
/// @notice Interface for ValueStore contracts
/// @author DIA (diadata.org)
interface IValueStore {
    /// @notice Get stored value for a key
    /// @param key The key to retrieve
    /// @return fairValue The stored fair value
    /// @return usdValue The stored USD value
    /// @return numerator The stored numerator
    /// @return denominator The stored denominator
    /// @return timestamp The timestamp of the update
    function getValue(
        string memory key
    )
        external
        view
        returns (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp);
}
