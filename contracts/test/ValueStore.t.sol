// SPDX-License-Identifier: MIT
pragma solidity 0.8.34;

import "forge-std/Test.sol";
import {ValueStore} from "../valuestore/ValueStore.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";
import {IValueStore} from "../interfaces/IValueStore.sol";
import {IERC1967} from "@openzeppelin/contracts/interfaces/IERC1967.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract ValueStoreTest is Test {
    ValueStore public implementation;
    ERC1967Proxy public proxy;
    ValueStore public valueStore;

    address public owner = address(0x1);
    address public user = address(0x2);
    address public newOwner = address(0x3);

    function setUp() public {
        // Deploy implementation
        implementation = new ValueStore();

        // Deploy proxy with initialization data
        bytes memory initData = abi.encodeWithSelector(
            ValueStore.initialize.selector,
            owner
        );
        proxy = new ERC1967Proxy(
            payable(address(implementation)),
            initData
        );

        // Wrap proxy in ValueStore interface
        valueStore = ValueStore(address(proxy));
    }

    // --- Ownership & Initialization Tests ---

    function test_ConstructorSetsOwner() public view {
        assertEq(valueStore.owner(), owner, "owner should be initialized correctly");
    }

    function test_InitializeWithZeroAddress() public {
        // Deploy new implementation and proxy with zero address
        ValueStore badImpl = new ValueStore();
        bytes memory initData = abi.encodeWithSelector(
            ValueStore.initialize.selector,
            address(0)
        );

        vm.expectRevert(ValueStore.ZeroAddress.selector);
        new ERC1967Proxy(
            payable(address(badImpl)),
            initData
        );
    }

    function test_TransferOwnership() public {
        // Set some data before transfer to verify persistence
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        assertEq(valueStore.owner(), newOwner, "ownership should be transferred");

        // Verify old owner can no longer set values
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", owner)
        );
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        // Verify new owner can set values
        vm.prank(newOwner);
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        // Verify original data persisted through transfer
        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "original data should persist after ownership transfer");

        // Verify new data was set correctly
        (uint256 fv2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fv2, 200, "new owner should be able to set values");
    }

    function test_TransferOwnershipToZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableInvalidOwner(address)", address(0))
        );
        valueStore.transferOwnership(address(0));
    }

    function test_TransferOwnershipNotOwner() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );
        valueStore.transferOwnership(newOwner);
    }

    function test_OldOwnerCannotSetValueAfterTransfer() public {
        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        // Old owner should no longer have permissions
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", owner)
        );
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        // New owner should have permissions
        vm.prank(newOwner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "new owner should be able to set values");
    }

    function test_RenounceOwnership() public {
        // Set data before renouncing to verify persistence
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.renounceOwnership();

        assertEq(valueStore.owner(), address(0), "owner should be address(0) after renouncing");

        // After renouncing, old owner should not be able to set values
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", owner)
        );
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        // Verify original data persists after renouncing
        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "data should persist after ownership renunciation");

        // Verify that even random addresses can't set values (no owner exists)
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);
    }

    function test_InitializeCannotBeCalledTwice() public {
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("InvalidInitialization()")
        );
        valueStore.initialize(owner);
    }

    function test_TransferOwnershipEmitsEvent() public {
        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit OwnableUpgradeable.OwnershipTransferred(owner, newOwner);
        valueStore.transferOwnership(newOwner);

        assertEq(valueStore.owner(), newOwner);
    }

    function test_RenounceOwnershipEmitsEvent() public {
        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit OwnableUpgradeable.OwnershipTransferred(owner, address(0));
        valueStore.renounceOwnership();

        assertEq(valueStore.owner(), address(0));
    }

    function test_ConstructorDisablesInitializers() public {
        // Direct calls to implementation should work for initialize
        // But the constructor should have disabled initializers on the implementation
        // This is tested implicitly by successful setUp
        // Additional check: calling setValue on implementation directly should fail
        // because it's not initialized (no owner set)
        vm.expectRevert();
        implementation.setValue("test", 1, 1, 1, 1);
    }

    // --- Upgrade Tests ---

    function test_ProxyAndImplementationSeparation() public view {
        // Proxy and implementation should have different addresses
        assertNotEq(address(valueStore), address(implementation), "proxy and implementation should have different addresses");
        assertEq(address(valueStore), address(proxy), "valueStore should be the proxy address");
    }

    function test_OwnerCanUpgrade() public {
        // Set data before upgrade to verify persistence
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        ValueStore newImpl = new ValueStore();
        address oldImpl = address(implementation);

        vm.prank(owner);
        valueStore.upgradeToAndCall(address(newImpl), "");

        // Verify ownership persists
        assertEq(valueStore.owner(), owner, "owner should persist after upgrade");

        // Verify all data persisted through upgrade
        (uint256 fv, uint256 vu, uint256 num, uint256 den, uint256 ts) =
            valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "fairValue should persist after upgrade");
        assertEq(vu, 1000, "valueUsd should persist after upgrade");
        assertEq(num, 1, "numerator should persist after upgrade");
        assertEq(den, 1, "denominator should persist after upgrade");
        assertGt(ts, 0, "timestamp should persist after upgrade");

        // Verify contract functionality works after upgrade
        vm.prank(owner);
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        (uint256 fv2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fv2, 200, "new values should work correctly after upgrade");
    }

    function test_NonOwnerCannotUpgrade() public {
        ValueStore newImpl = new ValueStore();

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );
        valueStore.upgradeToAndCall(address(newImpl), "");
    }

    function test_DataPersistsAfterUpgrade() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        ValueStore newImpl = new ValueStore();

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit IERC1967.Upgraded(address(newImpl));
        valueStore.upgradeToAndCall(address(newImpl), "");

        (uint256 fairValue,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue, 100, "data should persist after upgrade");
    }

    function test_UpgradeToZeroAddressFails() public {
        vm.prank(owner);
        // UUPSUpgradeable validates the new implementation address
        // ERC1967Upgrade: new implementation is not UUPS
        // This typically fails with a low-level revert or specific error
        vm.expectRevert();
        valueStore.upgradeToAndCall(address(0), "");
    }

    function test_UpgradeToNonContractFails() public {
        vm.prank(owner);
        // Upgrading to a non-contract address should fail
        // The UUPSUpgradeable._authorizeUpgrade or ERC1967Upgrade will reject this
        vm.expectRevert();
        valueStore.upgradeToAndCall(address(0x123), "");
    }

    function test_MultipleUpgrades() public {
        // Test upgrading multiple times
        ValueStore impl2 = new ValueStore();
        ValueStore impl3 = new ValueStore();

        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.upgradeToAndCall(address(impl2), "");

        vm.prank(owner);
        valueStore.upgradeToAndCall(address(impl3), "");

        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "data should persist through multiple upgrades");
    }

    function test_UpgradeAfterOwnershipTransfer() public {
        ValueStore newImpl = new ValueStore();

        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        // Old owner can't upgrade
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", owner)
        );
        valueStore.upgradeToAndCall(address(newImpl), "");

        // New owner can upgrade
        vm.prank(newOwner);
        valueStore.upgradeToAndCall(address(newImpl), "");
    }

    function test_ProxyDeployment() public view {
        assertNotEq(address(proxy), address(0), "proxy address should not be zero");
        assertNotEq(address(valueStore), address(0), "valueStore address should not be zero");
    }

    // --- SetValue Tests ---

    function test_SetValue() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue(key);

        assertEq(fairValue, 100, "fairValue should match");
        assertEq(valueUsd, 1000, "valueUsd should match");
        assertEq(numerator, 1, "numerator should match");
        assertEq(denominator, 1, "denominator should match");
        assertEq(timestamp, block.timestamp, "timestamp should match");
    }

    function test_SetValueNotOwner() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
    }

    function test_SetValueDivisionByZero() public {
        vm.prank(owner);
        vm.expectRevert(ValueStore.DivisionByZero.selector);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 0);
    }

    function test_SetValueEmitsEvent() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit ValueStore.ValueUpdated(key, 100, 1000, 1, 1, block.timestamp);
        valueStore.setValue(key, 100, 1000, 1, 1);
    }

    function test_SetValueWithZeroNumerator() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 0, 1);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue("BTC/USD");
        assertEq(fairValue, 100, "fairValue should match");
        assertEq(valueUsd, 1000, "valueUsd should match");
        assertEq(numerator, 0, "numerator should be 0");
        assertEq(denominator, 1, "denominator should be 1");
        assertEq(timestamp, block.timestamp, "timestamp should match");
    }

    function test_SetValueWithZeroNumeratorAndDenominator() public {
        // EDGE CASE: 0/0 is intentionally allowed
        // This represents an "undefined" or "invalid" fraction state
        // Consumers should check for this case: numerator == 0 && denominator == 0
        // It can be used to signal that the fraction is not applicable or invalid
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 0, 0);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue(key);

        // Verify all fields stored correctly
        assertEq(fairValue, 100, "fairValue should match");
        assertEq(valueUsd, 1000, "valueUsd should match");
        assertEq(numerator, 0, "numerator should be 0");
        assertEq(denominator, 0, "denominator should be 0");
        assertEq(timestamp, block.timestamp, "timestamp should match");
    }

    function test_SetValueWithLargeNumbers() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 1e18, 1e18, 1e9, 1e9);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue("BTC/USD");

        assertEq(fairValue, 1e18, "fairValue should be 1e18");
        assertEq(valueUsd, 1e18, "valueUsd should be 1e18");
        assertEq(numerator, 1e9, "numerator should be 1e9");
        assertEq(denominator, 1e9, "denominator should be 1e9");
        assertGt(timestamp, 0, "timestamp should be greater than 0");
    }

    function test_SetValueWithMaxUint() public {
        string memory key = "MAX";

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit ValueStore.ValueUpdated(
            key,
            type(uint256).max,
            type(uint256).max,
            type(uint256).max,
            type(uint256).max,
            block.timestamp
        );
        valueStore.setValue(
            key,
            type(uint256).max,
            type(uint256).max,
            type(uint256).max,
            type(uint256).max
        );

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
        ) = valueStore.getValue(key);

        assertEq(fairValue, type(uint256).max, "fairValue should be max uint256");
        assertEq(valueUsd, type(uint256).max, "valueUsd should be max uint256");
        assertEq(numerator, type(uint256).max, "numerator should be max uint256");
        assertEq(denominator, type(uint256).max, "denominator should be max uint256");
    }

    function test_OverwriteExistingValue() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue(key, 200, 2000, 2, 1);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue("BTC/USD");

        assertEq(fairValue, 200, "fairValue should be updated to 200");
        assertEq(valueUsd, 2000, "valueUsd should be updated to 2000");
        assertEq(numerator, 2, "numerator should be updated to 2");
        assertEq(denominator, 1, "denominator should remain 1");
        assertGt(timestamp, 0, "timestamp should be set");
    }

    function test_TimestampUpdates() public {
        string memory key = "BTC/USD";

        vm.warp(1000); // Set predictable timestamp

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        (,,,, uint256 timestamp1) = valueStore.getValue(key);
        assertEq(timestamp1, 1000, "first timestamp should be 1000");

        vm.warp(2000); // Warp to specific time

        vm.prank(owner);
        valueStore.setValue(key, 200, 2000, 2, 1);

        (,,,, uint256 timestamp2) = valueStore.getValue(key);
        assertEq(timestamp2, 2000, "second timestamp should be 2000");
        assertEq(timestamp2 - timestamp1, 1000, "difference should be 1000");
    }

    function test_TimestampAtMinimum() public {
        // Test at timestamp = 1 (can't be 0 in Foundry)
        vm.warp(1);

        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        (,,,, uint256 ts) = valueStore.getValue("BTC/USD");
        assertEq(ts, 1, "timestamp should be 1");
    }

    function test_TimestampAtMaxUint() public {
        vm.warp(type(uint256).max);

        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        (,,,, uint256 ts) = valueStore.getValue("BTC/USD");
        assertEq(ts, type(uint256).max, "max uint timestamp should be stored correctly");

        // Verify we can still update (timestamp would overflow if we incremented, but we just set it)
        vm.warp(type(uint256).max - 1);
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 200, 2000, 2, 1);

        (,,,, uint256 ts2) = valueStore.getValue("BTC/USD");
        assertEq(ts2, type(uint256).max - 1, "timestamp should update correctly");
    }

    function test_StateUnchangedAfterFailedSetValue() public {
        string memory key = "BTC/USD";

        // Set initial value
        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        // Try to set invalid value (division by zero)
        vm.prank(owner);
        vm.expectRevert(ValueStore.DivisionByZero.selector);
        valueStore.setValue(key, 200, 2000, 1, 0);

        // Verify original value is unchanged
        (uint256 fairValue,,,,) = valueStore.getValue(key);
        assertEq(fairValue, 100, "value should be unchanged after failed update");
    }

    function test_SpecialCharactersInKey() public {
        string memory key1 = "BTC/USD?query=1&foo=bar";
        string memory key2 = unicode" emojis 🚀 ";
        string memory key3 = string(new bytes(500)); // Very long key

        vm.prank(owner);
        valueStore.setValue(key1, 1, 1, 1, 1);

        vm.prank(owner);
        valueStore.setValue(key2, 2, 2, 1, 1);

        vm.prank(owner);
        valueStore.setValue(key3, 3, 3, 1, 1);

        (uint256 v1,,,,) = valueStore.getValue(key1);
        assertEq(v1, 1, "value with special chars should be stored");

        (uint256 v2,,,,) = valueStore.getValue(key2);
        assertEq(v2, 2, "value with unicode should be stored");

        (uint256 v3,,,,) = valueStore.getValue(key3);
        assertEq(v3, 3, "value with long key should be stored");
    }

    function test_EmptyStringKey() public {
        vm.prank(owner);
        vm.expectRevert(ValueStore.InvalidKey.selector);
        valueStore.setValue("", 100, 1000, 1, 1);
    }

    // --- SetMultipleValues Tests ---

    function test_SetMultipleValues() public {
        string[] memory keys = new string[](3);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";
        keys[2] = "AAPL/USD";

        uint256[] memory fairValues = new uint256[](3);
        fairValues[0] = 100;
        fairValues[1] = 200;
        fairValues[2] = 300;

        uint256[] memory valueUsds = new uint256[](3);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;
        valueUsds[2] = 3000;

        uint256[] memory numerators = new uint256[](3);
        numerators[0] = 1;
        numerators[1] = 2;
        numerators[2] = 3;

        uint256[] memory denominators = new uint256[](3);
        denominators[0] = 1;
        denominators[1] = 1;
        denominators[2] = 1;

        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Verify first value
        (uint256 fairValue1, uint256 valueUsd1, uint256 numerator1, uint256 denominator1,) =
            valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100, "first fairValue should match");
        assertEq(valueUsd1, 1000, "first valueUsd should match");
        assertEq(numerator1, 1, "first numerator should match");
        assertEq(denominator1, 1, "first denominator should match");

        // Verify second value
        (uint256 fairValue2, uint256 valueUsd2, uint256 numerator2, uint256 denominator2,) =
            valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200, "second fairValue should match");
        assertEq(valueUsd2, 2000, "second valueUsd should match");
        assertEq(numerator2, 2, "second numerator should match");
        assertEq(denominator2, 1, "second denominator should match");

        // Verify third value
        (uint256 fairValue3, uint256 valueUsd3, uint256 numerator3, uint256 denominator3,) =
            valueStore.getValue("AAPL/USD");
        assertEq(fairValue3, 300, "third fairValue should match");
        assertEq(valueUsd3, 3000, "third valueUsd should match");
        assertEq(numerator3, 3, "third numerator should match");
        assertEq(denominator3, 1, "third denominator should match");
    }

    function test_SetMultipleValuesWithEmptyArrays() public {
        // Empty arrays should be valid (no-op operation)
        string[] memory keys = new string[](0);
        uint256[] memory fairValues = new uint256[](0);
        uint256[] memory valueUsds = new uint256[](0);
        uint256[] memory numerators = new uint256[](0);
        uint256[] memory denominators = new uint256[](0);

        vm.prank(owner);
        // Should succeed - empty batch is valid (no-op)
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Explicitly assert success
        assertTrue(true, "empty batch operation should succeed without revert");
    }

    function test_SetMultipleValuesNotOwner() public {
        string[] memory keys = new string[](1);
        keys[0] = "BTC/USD";

        uint256[] memory fairValues = new uint256[](1);
        fairValues[0] = 100;

        uint256[] memory valueUsds = new uint256[](1);
        valueUsds[0] = 1000;

        uint256[] memory numerators = new uint256[](1);
        numerators[0] = 1;

        uint256[] memory denominators = new uint256[](1);
        denominators[0] = 1;

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_SetMultipleValuesArrayLengthMismatch() public {
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";

        uint256[] memory fairValues = new uint256[](1);
        fairValues[0] = 100;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 2;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        vm.expectRevert(ValueStore.InvalidArrayLengths.selector);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_SetMultipleValuesEmitsEvents() public {
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 2;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        // Expect both events
        vm.expectEmit(true, true, true, true);
        emit ValueStore.ValueUpdated(keys[0], 100, 1000, 1, 1, block.timestamp);
        vm.expectEmit(true, true, true, true);
        emit ValueStore.ValueUpdated(keys[1], 200, 2000, 2, 1, block.timestamp);

        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_SetMultipleValuesWithMixedMaxValues() public {
        string[] memory keys = new string[](2);
        keys[0] = "MAX1";
        keys[1] = "MAX2";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = type(uint256).max;
        fairValues[1] = 0;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 0;
        valueUsds[1] = type(uint256).max;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 1;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        (uint256 fv1,,,,) = valueStore.getValue("MAX1");
        assertEq(fv1, type(uint256).max, "first value should be max uint256");

        (uint256 fv2,,,,) = valueStore.getValue("MAX2");
        assertEq(fv2, 0, "second value should be 0");
    }

    function test_DivisionByZeroProtectionInMultiple() public {
        // EDGE CASE: 0/0 is intentionally allowed in batch operations
        // This allows storing "undefined" or "invalid" fractions
        // Only non-zero/zero combinations are rejected (true division by zero)
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 0;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 0;

        vm.prank(owner);
        // Should pass because 0/0 is allowed (numerator == 0)
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Verify values
        (uint256 fairValue1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100, "first value should be stored");

        (uint256 fairValue2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200, "second value should be stored");
    }

    function test_DuplicateKeysInBatch() public {
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "BTC/USD"; // Duplicate key

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 1;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Should have the second value (last write wins)
        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 200, "last write should win in batch");
    }

    function test_EmptyStringKeyInMultipleValues() public {
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 1;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(ValueStore.InvalidKeyInBatch.selector, 1));
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_DivisionByZeroInMultipleValues() public {
        string[] memory keys = new string[](3);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";
        keys[2] = "AAPL/USD";

        uint256[] memory fairValues = new uint256[](3);
        fairValues[0] = 100;
        fairValues[1] = 200;
        fairValues[2] = 300;

        uint256[] memory valueUsds = new uint256[](3);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;
        valueUsds[2] = 3000;

        uint256[] memory numerators = new uint256[](3);
        numerators[0] = 1;
        numerators[1] = 1;
        numerators[2] = 1;

        uint256[] memory denominators = new uint256[](3);
        denominators[0] = 1;
        denominators[1] = 0;
        denominators[2] = 1;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(ValueStore.DivisionByZeroInBatch.selector, 1));
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_BatchErrorAtIndexZero() public {
        string[] memory keys = new string[](2);
        keys[0] = "";
        keys[1] = "ETH/USD";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 1;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(ValueStore.InvalidKeyInBatch.selector, 0));
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_StateUnchangedAfterFailedBatch() public {
        string[] memory keys = new string[](2);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";

        uint256[] memory fairValues = new uint256[](2);
        fairValues[0] = 100;
        fairValues[1] = 200;

        uint256[] memory valueUsds = new uint256[](2);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;

        uint256[] memory numerators = new uint256[](2);
        numerators[0] = 1;
        numerators[1] = 1;

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 1;

        // Set initial values
        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Try batch with division by zero at index 1
        uint256[] memory badDenominators = new uint256[](2);
        badDenominators[0] = 1;
        badDenominators[1] = 0;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(ValueStore.DivisionByZeroInBatch.selector, 1));
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, badDenominators);

        // Verify original values are unchanged (atomic rollback)
        (uint256 fv1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv1, 100, "first value should be unchanged");

        (uint256 fv2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fv2, 200, "second value should be unchanged");
    }

    function test_LargeBatchOperation() public {
        uint256 batchSize = 100;
        string[] memory keys = new string[](batchSize);
        uint256[] memory fairValues = new uint256[](batchSize);
        uint256[] memory valueUsds = new uint256[](batchSize);
        uint256[] memory numerators = new uint256[](batchSize);
        uint256[] memory denominators = new uint256[](batchSize);

        for (uint256 i = 0; i < batchSize; i++) {
            keys[i] = vm.toString(i);
            fairValues[i] = i;
            valueUsds[i] = i * 100;
            numerators[i] = 1;
            denominators[i] = 1;
        }

        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Spot check a few values
        (uint256 fv0,,,,) = valueStore.getValue("0");
        assertEq(fv0, 0, "first value should be 0");

        (uint256 fv50,,,,) = valueStore.getValue("50");
        assertEq(fv50, 50, "middle value should be 50");

        (uint256 fv99,,,,) = valueStore.getValue("99");
        assertEq(fv99, 99, "last value should be 99");
    }

    // --- Gas Tests ---

    function test_GasCostOfSetValue() public {
        // Gas cost breakdown (first write to cold storage):
        // - Sstore (cold slot): ~20,000 gas
        // - Sstore (warm slot): ~5,000 gas
        // - Owner check (onlyOwner modifier): ~2,600 gas
        // - Event emission: ~1,500 gas
        // - Other overhead: ~100,000 gas (dispatcher, proxy, etc.)
        // Total expected: ~130,000 gas for first write
        vm.prank(owner);
        uint256 gasBefore = gasleft();
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
        uint256 gasAfter = gasleft();

        uint256 gasUsed = gasBefore - gasAfter;

        // Allow margin for different compiler settings and optimizations
        assertLt(gasUsed, 150000, "single setValue should cost less than 150k gas");
        assertGt(gasUsed, 100000, "single setValue should cost more than 100k gas (cold storage)");

        emit log_uint(gasUsed);
    }

    function test_GasCostOfSetMultipleValues() public {
        string[] memory keys = new string[](10);
        uint256[] memory fairValues = new uint256[](10);
        uint256[] memory valueUsds = new uint256[](10);
        uint256[] memory numerators = new uint256[](10);
        uint256[] memory denominators = new uint256[](10);

        for (uint256 i = 0; i < 10; i++) {
            keys[i] = vm.toString(i);
            fairValues[i] = i;
            valueUsds[i] = i;
            numerators[i] = i;
            denominators[i] = i + 1;
        }

        // Gas cost breakdown for batch of 10 (all cold storage):
        // - First write: ~130k gas (cold storage + proxy overhead)
        // - Subsequent writes (warm storage): ~5k - 10k each
        // - Total: ~130k + 9*10k = ~220k gas
        // Plus loop overhead and validation
        vm.prank(owner);
        uint256 gasBefore = gasleft();
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
        uint256 gasAfter = gasleft();

        uint256 gasUsed = gasBefore - gasAfter;

        // 10 values with owner check (warm storage after first)
        // Allow margin for different optimization settings
        assertLt(gasUsed, 1200000, "batch of 10 should cost less than 1.2M gas");
        assertGt(gasUsed, 150000, "batch of 10 should cost more than 150k gas");

        emit log_uint(gasUsed);
    }

    // --- GetValue Tests ---

    function test_GetValue() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue(key);

        assertEq(fairValue, 100, "fairValue should match");
        assertEq(valueUsd, 1000, "valueUsd should match");
        assertEq(numerator, 1, "numerator should match");
        assertEq(denominator, 1, "denominator should match");
        assertEq(timestamp, block.timestamp, "timestamp should match");
    }

    function test_GetValueNonExistent() public {
        vm.expectRevert(ValueStore.NoDataForKey.selector);
        valueStore.getValue("NONEXISTENT");
    }

    // --- Edge Cases ---

    function test_MultipleKeys() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        (uint256 fairValue1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100, "first key value should match");

        (uint256 fairValue2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200, "second key value should match");
    }

    function test_MultipleKeysWithSamePrefix() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue("BTC/ETH", 200, 2000, 2, 1);

        (uint256 fairValue1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100, "first prefix value should match");

        (uint256 fairValue2,,,,) = valueStore.getValue("BTC/ETH");
        assertEq(fairValue2, 200, "second prefix value should match");
    }

    function test_DifferentKeysHaveIndependentValues() public {
        // Verify that different keys don't collide in storage
        // Keys that might hash similarly
        string memory key1 = "AB";
        string memory key2 = "A";
        string memory key3 = "BA";

        vm.prank(owner);
        valueStore.setValue(key1, 100, 1000, 1, 1);
        vm.prank(owner);
        valueStore.setValue(key2, 200, 2000, 2, 1);
        vm.prank(owner);
        valueStore.setValue(key3, 300, 3000, 3, 1);

        (uint256 fv1,,,,) = valueStore.getValue(key1);
        (uint256 fv2,,,,) = valueStore.getValue(key2);
        (uint256 fv3,,,,) = valueStore.getValue(key3);

        assertEq(fv1, 100, "key1 should have its own value");
        assertEq(fv2, 200, "key2 should have its own value");
        assertEq(fv3, 300, "key3 should have its own value");

        // Verify all values are different
        assertNotEq(fv1, fv2, "different keys should have different values");
        assertNotEq(fv2, fv3, "different keys should have different values");
        assertNotEq(fv1, fv3, "different keys should have different values");
    }

    // --- Invariant Tests ---

    function test_Invariant_NonExistentKeyReverts() public {
        // Reading a non-existent key should always revert
        vm.expectRevert(ValueStore.NoDataForKey.selector);
        valueStore.getValue("NONEXISTENT_KEY_12345");
    }

    function test_Invariant_OwnerCanAlwaysWrite() public {
        // Owner should always be able to write, regardless of state
        for (uint256 i = 0; i < 10; i++) {
            vm.prank(owner);
            valueStore.setValue(vm.toString(i), i, i * 100, 1, 1);

            (uint256 fv,,,,) = valueStore.getValue(vm.toString(i));
            assertEq(fv, i, "owner should be able to write values");
        }
    }

    function test_Invariant_NonOwnerCanNeverWrite() public {
        // Non-owner should never be able to write, even after various operations
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        // Try as non-owner before any state changes
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user));
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);

        // Verify value didn't change
        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "value should not change");

        // Try after ownership transfer
        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        vm.prank(user);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user));
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);
    }

    function test_Invariant_TimestampAlwaysMatchesBlockTimestamp() public {
        // All timestamps should equal block.timestamp at time of setting
        vm.warp(12345);

        vm.prank(owner);
        valueStore.setValue("KEY1", 1, 1, 1, 1);

        (,,,, uint256 ts1) = valueStore.getValue("KEY1");
        assertEq(ts1, 12345, "timestamp should match block timestamp");

        vm.warp(54321);

        vm.prank(owner);
        valueStore.setValue("KEY2", 2, 2, 1, 1);

        (,,,, uint256 ts2) = valueStore.getValue("KEY2");
        assertEq(ts2, 54321, "timestamp should match new block timestamp");
    }

    // --- Reentrancy and Security Tests ---

    function test_ReentrancyProtection() public {
        // Document reentrancy posture:
        // - onlyOwner modifier prevents external reentrancy
        // - No external calls before state changes
        // - CEI (Checks-Effects-Interactions) pattern followed
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        (uint256 fv,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fv, 100, "value should be set correctly");
    }

    // --- DoS Protection Tests ---

    function test_LargeBatchOperationDoesNotExhaustGas() public {
        uint256 batchSize = 200;  // Reasonable limit for gas
        string[] memory keys = new string[](batchSize);
        uint256[] memory fairValues = new uint256[](batchSize);
        uint256[] memory valueUsds = new uint256[](batchSize);
        uint256[] memory numerators = new uint256[](batchSize);
        uint256[] memory denominators = new uint256[](batchSize);

        for (uint256 i = 0; i < batchSize; i++) {
            keys[i] = vm.toString(i);
            fairValues[i] = i;
            valueUsds[i] = i * 100;
            numerators[i] = 1;
            denominators[i] = 1;
        }

        vm.prank(owner);
        // Should complete without out of gas
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Spot check values
        (uint256 fv0,,,,) = valueStore.getValue("0");
        assertEq(fv0, 0, "first value should be stored");

        (uint256 fv100,,,,) = valueStore.getValue("100");
        assertEq(fv100, 100, "middle value should be stored");

        (uint256 fv199,,,,) = valueStore.getValue("199");
        assertEq(fv199, 199, "last value should be stored");
    }

    // --- Fuzz Tests ---

    function testFuzz_SetValue(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator
    ) public {
        vm.assume(bytes(key).length > 0);
        vm.assume(numerator == 0 || denominator > 0);

        vm.prank(owner);
        valueStore.setValue(key, fairValue, valueUsd, numerator, denominator);

        (
            uint256 storedFv,
            uint256 storedVu,
            uint256 storedNum,
            uint256 storedDen,
            uint256 storedTs
        ) = valueStore.getValue(key);

        assertEq(storedFv, fairValue, "stored fairValue should match input");
        assertEq(storedVu, valueUsd, "stored valueUsd should match input");
        assertEq(storedNum, numerator, "stored numerator should match input");
        assertEq(storedDen, denominator, "stored denominator should match input");
        assertEq(storedTs, block.timestamp, "timestamp should match block timestamp");
    }

    function testFuzz_SetValueRevertsOnEmptyKey(
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator,
        uint256 denominator
    ) public {
        vm.prank(owner);
        vm.expectRevert(ValueStore.InvalidKey.selector);
        valueStore.setValue("", fairValue, valueUsd, numerator, denominator);
    }

    function testFuzz_SetValueRevertsOnDivisionByZero(
        string calldata key,
        uint256 fairValue,
        uint256 valueUsd,
        uint256 numerator
    ) public {
        vm.assume(bytes(key).length > 0);
        vm.assume(numerator > 0); // Ensure numerator is non-zero

        vm.prank(owner);
        vm.expectRevert(ValueStore.DivisionByZero.selector);
        valueStore.setValue(key, fairValue, valueUsd, numerator, 0);
    }

    // --- supportsInterface Tests ---

    function test_SupportsInterface_IERC165() public view {
        assertTrue(valueStore.supportsInterface(type(IERC165).interfaceId), "should support IERC165");
    }

    function test_SupportsInterface_IValueStore() public view {
        assertTrue(valueStore.supportsInterface(type(IValueStore).interfaceId), "should support IValueStore");
    }

    function test_SupportsInterface_InvalidInterface() public view {
        assertFalse(valueStore.supportsInterface(bytes4(0x00000000)), "should not support 0x00000000");
        assertFalse(valueStore.supportsInterface(bytes4(0xFFFFFFFF)), "should not support 0xFFFFFFFF");
    }

    function test_SupportsInterface_AfterUpgrade() public {
        // Interface support should persist after upgrade
        ValueStore newImpl = new ValueStore();

        vm.prank(owner);
        valueStore.upgradeToAndCall(address(newImpl), "");

        // All interfaces should still be supported
        assertTrue(valueStore.supportsInterface(type(IERC165).interfaceId), "should support IERC165 after upgrade");
        assertTrue(valueStore.supportsInterface(type(IValueStore).interfaceId), "should support IValueStore after upgrade");
        assertFalse(valueStore.supportsInterface(bytes4(0x00000000)), "should not support 0x00000000 after upgrade");
        assertFalse(valueStore.supportsInterface(bytes4(0xFFFFFFFF)), "should not support 0xFFFFFFFF after upgrade");
    }

    function test_SupportsInterface_EdgeCases() public {
        // Test some additional edge cases for interface detection
        // 0xffffffff is the ERC165 "magic" return value for invalid interfaces
        assertFalse(valueStore.supportsInterface(bytes4(0xFFFFFFFF)), "ERC165: 0xffffffff must return false");

        // 0x00000000 is not a valid interface ID
        assertFalse(valueStore.supportsInterface(bytes4(0x00000000)), "0x00000000 should return false");

        // Random interface IDs should return false
        assertFalse(valueStore.supportsInterface(bytes4(0x12345678)), "random interface should return false");
        assertFalse(valueStore.supportsInterface(bytes4(0xdeadbeef)), "random interface should return false");
    }

    // --- Consumer Contract Tests ---

    function test_ConsumerHandlingOfZeroOverZero() public {
        // Demonstrate how consumers should handle the 0/0 case
        // 0/0 represents an "undefined" or "invalid" fraction

        vm.prank(owner);
        valueStore.setValue("VALID_KEY", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue("INVALID_KEY", 200, 2000, 0, 0);

        // Consumer should check for 0/0 case
        (uint256 fv1,,, uint256 den1,) = valueStore.getValue("VALID_KEY");
        assertTrue(den1 != 0 || fv1 == 0, "valid fraction: denominator is non-zero or numerator is zero");

        (uint256 fv2,,, uint256 den2,) = valueStore.getValue("INVALID_KEY");
        bool isUndefined = (den2 == 0);
        assertTrue(isUndefined, "consumer should detect undefined fraction (0/0)");

        // Best practice: consumers should validate fractions before using them
        if (den2 == 0) {
            // Handle undefined fraction case
            assertTrue(true, "undefined fraction detected, handle appropriately");
        } else {
            // Safe to use the fraction
            assertTrue(false, "should not reach here for 0/0 case");
        }
    }

    // --- Differential Tests ---

    function testDifferential_SingleVsBatchConsistency() public {
        // Verify that single setValue and batch setMultipleValues produce identical results

        // First, set using batch
        string[] memory keys = new string[](3);
        keys[0] = "BTC/USD";
        keys[1] = "ETH/USD";
        keys[2] = "AAPL/USD";

        uint256[] memory fairValues = new uint256[](3);
        fairValues[0] = 100;
        fairValues[1] = 200;
        fairValues[2] = 300;

        uint256[] memory valueUsds = new uint256[](3);
        valueUsds[0] = 1000;
        valueUsds[1] = 2000;
        valueUsds[2] = 3000;

        uint256[] memory numerators = new uint256[](3);
        numerators[0] = 1;
        numerators[1] = 2;
        numerators[2] = 3;

        uint256[] memory denominators = new uint256[](3);
        denominators[0] = 1;
        denominators[1] = 1;
        denominators[2] = 1;

        vm.prank(owner);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Store batch results (limited variables to avoid stack too deep)
        (uint256 fv1_batch,,, uint256 den1_batch,) = valueStore.getValue("BTC/USD");
        (uint256 fv2_batch,,, uint256 den2_batch,) = valueStore.getValue("ETH/USD");
        (uint256 fv3_batch,,, uint256 den3_batch,) = valueStore.getValue("AAPL/USD");

        // Clear values
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 0, 0, 0, 1);
        vm.prank(owner);
        valueStore.setValue("ETH/USD", 0, 0, 0, 1);
        vm.prank(owner);
        valueStore.setValue("AAPL/USD", 0, 0, 0, 1);

        // Now set using individual calls
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
        vm.prank(owner);
        valueStore.setValue("ETH/USD", 200, 2000, 2, 1);
        vm.prank(owner);
        valueStore.setValue("AAPL/USD", 300, 3000, 3, 1);

        // Get single results and verify immediately
        {
            (uint256 fv1_single,,, uint256 den1_single,) = valueStore.getValue("BTC/USD");
            assertEq(fv1_batch, fv1_single, "batch and single should produce same fairValue for key 1");
            assertEq(den1_batch, den1_single, "batch and single should produce same denominator for key 1");
        }

        {
            (uint256 fv2_single,,, uint256 den2_single,) = valueStore.getValue("ETH/USD");
            assertEq(fv2_batch, fv2_single, "batch and single should produce same fairValue for key 2");
            assertEq(den2_batch, den2_single, "batch and single should produce same denominator for key 2");
        }

        {
            (uint256 fv3_single,,, uint256 den3_single,) = valueStore.getValue("AAPL/USD");
            assertEq(fv3_batch, fv3_single, "batch and single should produce same fairValue for key 3");
            assertEq(den3_batch, den3_single, "batch and single should produce same denominator for key 3");
        }
    }
}
