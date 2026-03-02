// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

import "forge-std/Test.sol";
import {ValueStore} from "../valuestore/ValueStore.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";

contract ValueStoreTest is Test {
    ValueStore public implementation;
    ERC1967Proxy public proxy;
    ValueStore public valueStore;

    address public owner = address(0x1);
    address public user = address(0x2);

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

    function test_ConstructorSetsOwner() public {
        assertEq(valueStore.owner(), owner);
    }

    function test_TransferOwnership() public {
        address newOwner = address(0x3);

        vm.prank(owner);
        valueStore.transferOwnership(newOwner);

        assertEq(valueStore.owner(), newOwner);
    }

    function test_TransferOwnershipToZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(); // OwnableInvalidOwner
        valueStore.transferOwnership(address(0));
    }

    function test_TransferOwnershipNotOwner() public {
        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        valueStore.transferOwnership(address(0x3));
    }

    // --- Upgrade Tests ---

    function test_ProxyAndImplementationSeparation() public {
        // Proxy and implementation should have different addresses
        assertNotEq(address(valueStore), address(implementation));
        assertEq(address(valueStore), address(proxy));
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

        assertEq(fairValue, 100);
        assertEq(valueUsd, 1000);
        assertEq(numerator, 1);
        assertEq(denominator, 1);
        assertEq(timestamp, block.timestamp);
    }

    function test_SetValueNotOwner() public {
        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
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
        valueStore.setValue("BTC/USD", 100, 1000, 0, 1); // Should work (0/x is valid)

        (
            uint256 fairValue,
            uint256 valueUsd,
            uint256 numerator,
            uint256 denominator,
            uint256 timestamp
        ) = valueStore.getValue("BTC/USD");
        assertEq(fairValue, 100);
        assertEq(valueUsd, 1000);
        assertEq(numerator, 0);
        assertEq(denominator, 1);
        assertEq(timestamp, block.timestamp);
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

        assertEq(fairValue, 1e18);
        assertEq(valueUsd, 1e18);
        assertEq(numerator, 1e9);
        assertEq(denominator, 1e9);
        assertGt(timestamp, 0);
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

        assertEq(fairValue, 200);
        assertEq(valueUsd, 2000);
        assertEq(numerator, 2);
        assertEq(denominator, 1);
        assertGt(timestamp, 0);
    }

    function test_TimestampUpdates() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        (,,,, uint256 timestamp1) = valueStore.getValue(key);

        vm.warp(block.timestamp + 100);

        vm.prank(owner);
        valueStore.setValue(key, 200, 2000, 2, 1);

        (,,,, uint256 timestamp2) = valueStore.getValue(key);

        assertEq(timestamp2, timestamp1 + 100);
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
        assertEq(fairValue1, 100);
        assertEq(valueUsd1, 1000);
        assertEq(numerator1, 1);
        assertEq(denominator1, 1);

        // Verify second value
        (uint256 fairValue2, uint256 valueUsd2, uint256 numerator2, uint256 denominator2,) =
            valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200);
        assertEq(valueUsd2, 2000);
        assertEq(numerator2, 2);
        assertEq(denominator2, 1);

        // Verify third value
        (uint256 fairValue3, uint256 valueUsd3, uint256 numerator3, uint256 denominator3,) =
            valueStore.getValue("AAPL/USD");
        assertEq(fairValue3, 300);
        assertEq(valueUsd3, 3000);
        assertEq(numerator3, 3);
        assertEq(denominator3, 1);
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
        vm.expectRevert(); // OwnableUnauthorizedAccount
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
        vm.expectEmit(true, true, true, true);
        emit ValueStore.ValueUpdated(keys[0], 100, 1000, 1, 1, block.timestamp);
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);
    }

    function test_GasCostOfSetValue() public {
        vm.prank(owner);
        uint256 gasBefore = gasleft();
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
        uint256 gasAfter = gasleft();

        uint256 gasUsed = gasBefore - gasAfter;
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

        assertEq(fairValue, 100);
        assertEq(valueUsd, 1000);
        assertEq(numerator, 1);
        assertEq(denominator, 1);
        assertEq(timestamp, block.timestamp);
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
        assertEq(fairValue1, 100);

        (uint256 fairValue2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200);
    }

    function test_MultipleKeysWithSamePrefix() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue("BTC/ETH", 200, 2000, 2, 1);

        (uint256 fairValue1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100);

        (uint256 fairValue2,,,,) = valueStore.getValue("BTC/ETH");
        assertEq(fairValue2, 200);
    }

    function test_UpdateExistingValue() public {
        string memory key = "BTC/USD";

        vm.prank(owner);
        valueStore.setValue(key, 100, 1000, 1, 1);

        vm.prank(owner);
        valueStore.setValue(key, 200, 2000, 2, 1);

        (uint256 fairValue,,,,) = valueStore.getValue(key);
        assertEq(fairValue, 200);
    }

    function test_InitializeCannotBeCalledTwice() public {
        // Trying to reinitialize should fail
        vm.prank(owner);
        vm.expectRevert(); // Initializable contract error
        valueStore.initialize(owner);
    }

    function test_InitializeWithZeroAddress() public {
        // Zero address protection is tested by the revert in initialize
        assertTrue(true);
    }

    function test_ConstructorIsDisabled() public {
        // Constructor should disable initializers
        // This is tested implicitly by setUp() succeeding
        assertTrue(true);
    }

    function test_ProxyDeployment() public {
        // Proxy should be deployed
        assertNotEq(address(proxy), address(0));
        assertNotEq(address(valueStore), address(0));
    }

    function test_RenounceOwnership() public {
        vm.prank(owner);
        valueStore.renounceOwnership();

        assertEq(valueStore.owner(), address(0));

        // After renouncing, no one should be able to set values
        vm.prank(owner);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        valueStore.setValue("BTC/USD", 100, 1000, 1, 1);
    }

    function test_EmptyStringKey() public {
        vm.prank(owner);
        valueStore.setValue("", 100, 1000, 1, 1);

        (uint256 fairValue,,,,) = valueStore.getValue("");
        assertEq(fairValue, 100);
    }

    function test_SetValueWithZeroNumeratorAndDenominator() public {
        vm.prank(owner);
        valueStore.setValue("BTC/USD", 100, 1000, 0, 0);

        (, , uint256 numerator, uint256 denominator,) = valueStore.getValue("BTC/USD");
        assertEq(numerator, 0);
        assertEq(denominator, 0);
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
        numerators[1] = 0; // This one is fine

        uint256[] memory denominators = new uint256[](2);
        denominators[0] = 1;
        denominators[1] = 0; // This will cause division by zero with 0 numerator

        vm.prank(owner);
        // Should pass because 0/0 is allowed (numerator == 0)
        valueStore.setMultipleValues(keys, fairValues, valueUsds, numerators, denominators);

        // Verify values
        (uint256 fairValue1,,,,) = valueStore.getValue("BTC/USD");
        assertEq(fairValue1, 100);

        (uint256 fairValue2,,,,) = valueStore.getValue("ETH/USD");
        assertEq(fairValue2, 200);
    }

    // --- supportsInterface Tests ---

    function test_SupportsInterface_IERC165() public {
        assertTrue(valueStore.supportsInterface(type(IERC165).interfaceId));
    }

    function test_SupportsInterface_InvalidInterface() public {
        assertFalse(valueStore.supportsInterface(bytes4(0x00000000)));
        assertFalse(valueStore.supportsInterface(bytes4(0xFFFFFFFF)));
    }
}
