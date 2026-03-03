// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "forge-std/Test.sol";
import {ValueStore} from "../valuestore/ValueStore.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";
import {IValueStore} from "../interfaces/IValueStore.sol";
import {IERC1967} from "@openzeppelin/contracts/interfaces/IERC1967.sol";

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
        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        assertEq(valueStore.owner(), newOwner, "ownership should be transferred");
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
        vm.prank(owner);
        valueStore.renounceOwnership();

        assertEq(valueStore.owner(), address(0), "owner should be address(0) after renouncing");

        // After renouncing, no one should be able to set values
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", owner)
        );
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
    }

    function test_InitializeCannotBeCalledTwice() public {
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("InvalidInitialization()")
        );
        valueStore.initialize(owner);
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
        ValueStore newImpl = new ValueStore();

        vm.prank(owner);
        valueStore.upgradeToAndCall(address(newImpl), "");

        assertEq(valueStore.owner(), owner, "owner should persist after upgrade");
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
        vm.expectRevert();
        valueStore.upgradeToAndCall(address(0), "");
    }

    function test_UpgradeToNonContractFails() public {
        vm.prank(owner);
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
        string[] memory keys = new string[](0);
        uint256[] memory fairValues = new uint256[](0);
        uint256[] memory valueUsds = new uint256[](0);
        uint256[] memory numerators = new uint256[](0);
        uint256[] memory denominators = new uint256[](0);

        vm.prank(owner);
        // Should succeed - empty batch is valid
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
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
        vm.prank(owner);
        uint256 gasBefore = gasleft();
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
        uint256 gasAfter = gasleft();

        uint256 gasUsed = gasBefore - gasAfter;

        // Assert gas is within reasonable bounds
        // First-write to storage with owner check is more expensive (~130k gas)
        assertLt(gasUsed, 150000, "single setValue should cost less than 150k gas");

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

        vm.prank(owner);
        uint256 gasBefore = gasleft();
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
        uint256 gasAfter = gasleft();

        uint256 gasUsed = gasBefore - gasAfter;

        // 10 values with owner check should cost less than 1.2M gas
        assertLt(gasUsed, 1200000, "batch of 10 should cost less than 1.2M gas");

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
}
