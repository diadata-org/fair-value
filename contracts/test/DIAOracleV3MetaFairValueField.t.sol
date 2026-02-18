// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.30;

import "forge-std/Test.sol";

// Import the contracts
import {DIAOracleV3MetaFairValueField} from "../metacontract/DIAOracleV3MetaFairValueField.sol";

// Mock ValueStore contract for testing
contract MockValueStore {
    mapping(string => uint256) public fairValues;
    mapping(string => uint256) public usdValues;
    mapping(string => uint256) public numerators;
    mapping(string => uint256) public denominators;
    mapping(string => uint256) public timestamps;

    function setValue(
        string memory key,
        uint256 fairValue,
        uint256 usdValue,
        uint256 numerator,
        uint256 denominator,
        uint256 timestamp
    ) external {
        fairValues[key] = fairValue;
        usdValues[key] = usdValue;
        numerators[key] = numerator;
        denominators[key] = denominator;
        timestamps[key] = timestamp;
    }

    function getValue(string memory key)
        external
        view
        returns (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp)
    {
        // Revert if no data set for this key (mimics real ValueStore behavior)
        if (timestamps[key] == 0) {
            revert("No data for key");
        }
        return (fairValues[key], usdValues[key], numerators[key], denominators[key], timestamps[key]);
    }
}

// Test data structures
struct OracleData {
    uint256 fairValue;
    uint256 usdValue;
    uint256 numerator;
    uint256 denominator;
    uint256 timestamp;
    bool shouldRevert; // If true, getValue will revert
}

struct TestCase {
    string name;
    OracleData[] oracleValues;
    uint256 threshold;
    uint256 timeoutSeconds;
    uint256 blockTimestamp;
    uint256 expectedFairValue;
    uint256 expectedUsdValue;
    uint256 expectedNumerator;
    uint256 expectedDenominator;
    bool shouldRevert;
    bytes4 expectedError;  
    uint256 expectedErrorCount;  
}

 contract BaseTest is Test {
    DIAOracleV3MetaFairValueField public oracle;
    MockValueStore[] public stores;

    address public owner = address(0x1);
    address public user = address(0x2);

    string constant TEST_KEY = "BTC/USD";

    function setUp() public virtual {
        vm.prank(owner);
        oracle = new DIAOracleV3MetaFairValueField(owner);
    }

    function createMockStore() internal returns (MockValueStore) {
        MockValueStore store = new MockValueStore();
        stores.push(store);
        vm.prank(owner);
        oracle.addValueStore(address(store));
        return store;
    }

    function setStoreValues(
        MockValueStore store,
        string memory key,
        uint256 fairValue,
        uint256 usdValue,
        uint256 numerator,
        uint256 denominator,
        uint256 timestamp
    ) internal {
        if (timestamp == 0) {
             return;
        }
        store.setValue(key, fairValue, usdValue, numerator, denominator, timestamp);
    }

    function runTestCase(TestCase memory testCase) internal {
        // Reset oracle state
        uint256 numStores = stores.length;
        for (uint256 i = 0; i < numStores; i++) {
            vm.prank(owner);
            oracle.removeValueStore(address(stores[i]));
        }
        delete stores;

        // Execute test
        if (testCase.shouldRevert) {
            if (testCase.expectedError == DIAOracleV3MetaFairValueField.ThresholdNotMet.selector) {
                // Configure oracle first
                vm.prank(owner);
                oracle.setThreshold(testCase.threshold);
                vm.prank(owner);
                oracle.setTimeoutSeconds(testCase.timeoutSeconds);

                // Set block timestamp
                if (testCase.blockTimestamp > 0) {
                    vm.warp(testCase.blockTimestamp);
                }

                // Create stores and set values
                for (uint256 i = 0; i < testCase.oracleValues.length; i++) {
                    MockValueStore store = createMockStore();
                    OracleData memory data = testCase.oracleValues[i];

                    if (!data.shouldRevert) {
                        setStoreValues(
                            store,
                            TEST_KEY,
                            data.fairValue,
                            data.usdValue,
                            data.numerator,
                            data.denominator,
                            data.timestamp
                        );
                    }
                }

                // Then expect revert
                vm.expectRevert(
                    abi.encodeWithSelector(
                        DIAOracleV3MetaFairValueField.ThresholdNotMet.selector,
                        testCase.expectedErrorCount,
                        testCase.threshold
                    )
                );
                oracle.getMedianValues(TEST_KEY);
            } else if (testCase.expectedError == DIAOracleV3MetaFairValueField.InvalidThreshold.selector) {
                // Expect revert in setThreshold itself
                vm.prank(owner);
                vm.expectRevert(
                    abi.encodeWithSelector(
                        DIAOracleV3MetaFairValueField.InvalidThreshold.selector,
                        0
                    )
                );
                oracle.setThreshold(testCase.threshold);
            } else if (testCase.expectedError == DIAOracleV3MetaFairValueField.InvalidTimeOut.selector) {
                // Need to set threshold first, then expect revert in setTimeoutSeconds
                vm.prank(owner);
                oracle.setThreshold(testCase.threshold > 0 ? testCase.threshold : 1);

                vm.prank(owner);
                vm.expectRevert(
                    abi.encodeWithSelector(
                        DIAOracleV3MetaFairValueField.InvalidTimeOut.selector,
                        0
                    )
                );
                oracle.setTimeoutSeconds(testCase.timeoutSeconds);
            }
        } else {
            // Configure oracle
            vm.prank(owner);
            oracle.setThreshold(testCase.threshold);
            vm.prank(owner);
            oracle.setTimeoutSeconds(testCase.timeoutSeconds);

            // Set block timestamp
            if (testCase.blockTimestamp > 0) {
                vm.warp(testCase.blockTimestamp);
            }

            // Create stores and set values
            for (uint256 i = 0; i < testCase.oracleValues.length; i++) {
                MockValueStore store = createMockStore();
                OracleData memory data = testCase.oracleValues[i];

                if (!data.shouldRevert) {
                    setStoreValues(
                        store,
                        TEST_KEY,
                        data.fairValue,
                        data.usdValue,
                        data.numerator,
                        data.denominator,
                        data.timestamp
                    );
                }
            }

            DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

            // Assert results
            assertEq(
                result.fairValue,
                testCase.expectedFairValue,
                string(abi.encodePacked(testCase.name, ": fairValue mismatch"))
            );
            assertEq(
                result.usdValue,
                testCase.expectedUsdValue,
                string(abi.encodePacked(testCase.name, ": usdValue mismatch"))
            );
            assertEq(
                result.numerator,
                testCase.expectedNumerator,
                string(abi.encodePacked(testCase.name, ": numerator mismatch"))
            );
            assertEq(
                result.denominator,
                testCase.expectedDenominator,
                string(abi.encodePacked(testCase.name, ": denominator mismatch"))
            );
        }
    }
}

 
contract GetMedianValuesHappyPathTest is BaseTest {
    function test_HappyPathScenarios() public {
         TestCase[] memory testCases = new TestCase[](4);

        // Test Case 1: Odd number of oracles (3 stores)
        testCases[0] = TestCase({
            name: "Odd number of oracles",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 200,  // Median of [100, 200, 300]
            expectedUsdValue: 2000,
            expectedNumerator: 2,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[0].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[0].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[0].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000, false);

        // Test Case 2: Even number of oracles (4 stores)
        testCases[1] = TestCase({
            name: "Even number of oracles",
            oracleValues: new OracleData[](4),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 250,  // Average of [100, 200, 300, 400] middle = (200 + 300) / 2
            expectedUsdValue: 2500,
            expectedNumerator: 2,    // Average of 2 and 3
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[1].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[1].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[1].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000, false);
        testCases[1].oracleValues[3] = OracleData(400, 4000, 4, 1, 1000000, false);

        // Test Case 3: All same values
        testCases[2] = TestCase({
            name: "All same values",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 150,
            expectedUsdValue: 1500,
            expectedNumerator: 3,
            expectedDenominator: 2,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[2].oracleValues[0] = OracleData(150, 1500, 3, 2, 1000000, false);
        testCases[2].oracleValues[1] = OracleData(150, 1500, 3, 2, 1000000, false);
        testCases[2].oracleValues[2] = OracleData(150, 1500, 3, 2, 1000000, false);

        // Test Case 4: Five oracles
        testCases[3] = TestCase({
            name: "Five oracles",
            oracleValues: new OracleData[](5),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 300,  // Median of [100, 200, 300, 400, 500]
            expectedUsdValue: 3000,
            expectedNumerator: 3,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[3].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[3].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[3].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000, false);
        testCases[3].oracleValues[3] = OracleData(400, 4000, 4, 1, 1000000, false);
        testCases[3].oracleValues[4] = OracleData(500, 5000, 5, 1, 1000000, false);

        // Run all test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            runTestCase(testCases[i]);
        }
    }
}

 
contract GetMedianValuesThresholdTest is BaseTest {
    function test_ThresholdScenarios() public {
        TestCase[] memory testCases = new TestCase[](3);

        // Test Case 1: Meets minimum threshold
        testCases[0] = TestCase({
            name: "Meets minimum threshold",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 200,
            expectedUsdValue: 2000,
            expectedNumerator: 2,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[0].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[0].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[0].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000, false);

        // Test Case 2: Threshold not met
        testCases[1] = TestCase({
            name: "Threshold not met",
            oracleValues: new OracleData[](1),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 0,
            expectedUsdValue: 0,
            expectedNumerator: 0,
            expectedDenominator: 0,
            shouldRevert: true,
            expectedError: DIAOracleV3MetaFairValueField.ThresholdNotMet.selector,
            expectedErrorCount: 1  // Only 1 valid oracle
        });
        testCases[1].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);

        // Test Case 3: Threshold set to zero (should revert in setThreshold)
        testCases[2] = TestCase({
            name: "Invalid threshold zero",
            oracleValues: new OracleData[](0),
            threshold: 0,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 0,
            expectedUsdValue: 0,
            expectedNumerator: 0,
            expectedDenominator: 0,
            shouldRevert: true,
            expectedError: DIAOracleV3MetaFairValueField.InvalidThreshold.selector,
            expectedErrorCount: 0
        });

        // Run all test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            runTestCase(testCases[i]);
        }
    }
}

 
contract GetMedianValuesTimeoutTest is BaseTest {
    function test_TimeoutScenarios() public {
        TestCase[] memory testCases = new TestCase[](4);

        // Test Case 1: Filters out stale data
        testCases[0] = TestCase({
            name: "Filters out stale data",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 150,  // Median of [100, 200] (300 is stale)
            expectedUsdValue: 1500,
            expectedNumerator: 1,    // Average of 1 and 2
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[0].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);       // Fresh
        testCases[0].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);       // Fresh
        testCases[0].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000 - 3700, false); // Stale (> 3600)

        // Test Case 2: All data stale
        testCases[1] = TestCase({
            name: "All data stale",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 0,
            expectedUsdValue: 0,
            expectedNumerator: 0,
            expectedDenominator: 0,
            shouldRevert: true,
            expectedError: DIAOracleV3MetaFairValueField.ThresholdNotMet.selector,
            expectedErrorCount: 0  // No fresh data
        });
        testCases[1].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000 - 3700, false);
        testCases[1].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000 - 3700, false);
        testCases[1].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000 - 3700, false);

        // Test Case 3: Timeout set to zero
        testCases[2] = TestCase({
            name: "Invalid timeout zero",
            oracleValues: new OracleData[](0),
            threshold: 2,
            timeoutSeconds: 0,
            blockTimestamp: 1000000,
            expectedFairValue: 0,
            expectedUsdValue: 0,
            expectedNumerator: 0,
            expectedDenominator: 0,
            shouldRevert: true,
            expectedError: DIAOracleV3MetaFairValueField.InvalidTimeOut.selector,
            expectedErrorCount: 0
        });

        // Test Case 4: Boundary timeout
        testCases[3] = TestCase({
            name: "Boundary timeout",
            oracleValues: new OracleData[](2),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 150,  // Average of [100, 200]
            expectedUsdValue: 1500,
            expectedNumerator: 1,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[3].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000 - 3600, false); // Exactly at boundary
        testCases[3].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);

        // Run all test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            runTestCase(testCases[i]);
        }
    }
}

 
contract GetMedianValuesFailureHandlingTest is BaseTest {
    function test_FailureHandlingScenarios() public {
        TestCase[] memory testCases = new TestCase[](2);

        // Test Case 1: Handles failing oracle calls
        testCases[0] = TestCase({
            name: "Handles failing oracle calls",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 150,  // Median of [100, 200] (3rd reverts)
            expectedUsdValue: 1500,
            expectedNumerator: 1,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[0].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[0].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[0].oracleValues[2] = OracleData(0, 0, 0, 0, 0, true); // Will revert

        // Test Case 2: Most oracles failing
        testCases[1] = TestCase({
            name: "Most oracles failing",
            oracleValues: new OracleData[](5),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 150,  // Median of [100, 200]
            expectedUsdValue: 1500,
            expectedNumerator: 1,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[1].oracleValues[0] = OracleData(0, 0, 0, 0, 0, true); // Reverts
        testCases[1].oracleValues[1] = OracleData(0, 0, 0, 0, 0, true); // Reverts
        testCases[1].oracleValues[2] = OracleData(0, 0, 0, 0, 0, true); // Reverts
        testCases[1].oracleValues[3] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[1].oracleValues[4] = OracleData(200, 2000, 2, 1, 1000000, false);

        // Run all test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            runTestCase(testCases[i]);
        }
    }
}

 
contract GetMedianValuesEdgeCaseTest is BaseTest {
    function test_EdgeCaseScenarios() public {
        TestCase[] memory testCases = new TestCase[](3);

        // Test Case 1: Large numbers
        testCases[0] = TestCase({
            name: "Large numbers",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 2e18,  // Median of [1e18, 2e18, 3e18]
            expectedUsdValue: 2e18,
            expectedNumerator: 2,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[0].oracleValues[0] = OracleData(1e18, 1e18, 1, 1, 1000000, false);
        testCases[0].oracleValues[1] = OracleData(2e18, 2e18, 2, 1, 1000000, false);
        testCases[0].oracleValues[2] = OracleData(3e18, 3e18, 3, 1, 1000000, false);

        // Test Case 2: Zero values
        testCases[1] = TestCase({
            name: "Zero values",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 0,  // Median of [0, 0, 100]
            expectedUsdValue: 0,
            expectedNumerator: 0,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[1].oracleValues[0] = OracleData(0, 0, 0, 1, 1000000, false);
        testCases[1].oracleValues[1] = OracleData(0, 0, 0, 1, 1000000, false);
        testCases[1].oracleValues[2] = OracleData(100, 1000, 1, 1, 1000000, false);

        // Test Case 3: Different keys (uses different key)
        testCases[2] = TestCase({
            name: "Different keys - BTC",
            oracleValues: new OracleData[](3),
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 200,
            expectedUsdValue: 2000,
            expectedNumerator: 2,
            expectedDenominator: 1,
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });
        testCases[2].oracleValues[0] = OracleData(100, 1000, 1, 1, 1000000, false);
        testCases[2].oracleValues[1] = OracleData(200, 2000, 2, 1, 1000000, false);
        testCases[2].oracleValues[2] = OracleData(300, 3000, 3, 1, 1000000, false);

        // Run all test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            // For test case 3, we test with a different key
            if (i == 2) {
                // Create stores with both BTC and ETH data
                delete stores;
                vm.prank(owner);
                oracle = new DIAOracleV3MetaFairValueField(owner);
                vm.prank(owner);
                oracle.setThreshold(testCases[i].threshold);
                vm.prank(owner);
                oracle.setTimeoutSeconds(testCases[i].timeoutSeconds);
                vm.warp(testCases[i].blockTimestamp);

                for (uint256 j = 0; j < 3; j++) {
                    MockValueStore store = createMockStore();
                    uint256 ethValue = 2000 + (j * 100);
                    store.setValue("ETH/USD", ethValue, ethValue * 10, j + 1, 1, 1000000);
                    setStoreValues(store, "BTC/USD",
                        testCases[i].oracleValues[j].fairValue,
                        testCases[i].oracleValues[j].usdValue,
                        testCases[i].oracleValues[j].numerator,
                        testCases[i].oracleValues[j].denominator,
                        testCases[i].oracleValues[j].timestamp
                    );
                }

                DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues("BTC/USD");
                assertEq(result.fairValue, testCases[i].expectedFairValue);
            } else {
                runTestCase(testCases[i]);
            }
        }
    }
}


contract NumeratorDenominatorTest is BaseTest {
    function setUp() public override {
        super.setUp();
        // Initialize threshold and timeout for these tests
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
    }

    function test_NumeratorDenominator_OddOracles() public {
        // Setup: Create 3 stores with different numerators/denominators
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Store 1: 1/2 = 0.5
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 2, timestamp);
        // Store 2: 2/3 = 0.666... (should be median)
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 3, timestamp);
        // Store 3: 3/4 = 0.75
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 4, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Median should be the middle value (store2)
        assertEq(result.numerator, 2, "Numerator should be 2");
        assertEq(result.denominator, 3, "Denominator should be 3");
    }

    function test_NumeratorDenominator_EvenOracles() public {
        // Setup: Create 4 stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Store 1: 1/2
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 2, timestamp);
        // Store 2: 2/4 = 0.5
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 4, timestamp);
        // Store 3: 3/6 = 0.5
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 6, timestamp);
        // Store 4: 4/8 = 0.5
        setStoreValues(store4, TEST_KEY, 400, 4000, 4, 8, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Average of middle two: (2/4 + 3/6) / 2 = (2 + 3) / 2 / (4 + 6) / 2
        // But the contract calculates: (2 + 3) / 2 = 2, (4 + 6) / 2 = 5
        assertEq(result.numerator, 2, "Numerator should be average of 2 and 3");
        assertEq(result.denominator, 5, "Denominator should be average of 4 and 6");
    }

    function test_NumeratorDenominator_AllZero() public {
        // Test with all zero numerators
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 0, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 0, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 0, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, 0, "Numerator should be 0");
        assertEq(result.denominator, 1, "Denominator should be 1");
    }

    function test_NumeratorDenominator_LargeValues() public {
        // Test with large numerator/denominator values
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1000000, 2000000, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2000000, 3000000, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 3000000, 4000000, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, 2000000, "Numerator should be 2000000");
        assertEq(result.denominator, 3000000, "Denominator should be 3000000");
    }

    function test_NumeratorDenominator_SameDenominatorDifferentNumerators() public {
        // Test same denominator, different numerators
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 100, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 5, 100, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 9, 100, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, 5, "Numerator should be 5 (median)");
        assertEq(result.denominator, 100, "Denominator should be 100");
    }

    function test_NumeratorDenominator_SameNumeratorDifferentDenominators() public {
        // Test same numerator, different denominators
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 100, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 100, 5, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 100, 9, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, 100, "Numerator should be 100");
        assertEq(result.denominator, 5, "Denominator should be 5 (median)");
    }

    function test_NumeratorDenominator_VaryingFractions() public {
        // Test with varying fractions
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();
        MockValueStore store5 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);     // 1.0
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 2, timestamp);     // 0.5
        setStoreValues(store3, TEST_KEY, 300, 3000, 2, 3, timestamp);     // 0.666...
        setStoreValues(store4, TEST_KEY, 400, 4000, 3, 4, timestamp);     // 0.75
        setStoreValues(store5, TEST_KEY, 500, 5000, 1, 10, timestamp);    // 0.1

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorting is done by usdValues 
        // Sorted by fairValue: [100, 200, 300, 400, 500]
        // Corresponding numerators: [1, 1, 2, 3, 1]
        // Corresponding denominators: [1, 2, 3, 4, 10]
        // Median (3rd element, fairValue=300): numerator=2, denominator=3
        assertEq(result.numerator, 2, "Numerator should be 2");
        assertEq(result.denominator, 3, "Denominator should be 3");
    }

    function test_NumeratorDenominator_WithStaleData() public {
        // Test that stale data is excluded from numerator/denominator calculation
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 currentTimestamp = 1000000;
        vm.warp(currentTimestamp);

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 2, currentTimestamp);           // Fresh
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 3, currentTimestamp);           // Fresh
        setStoreValues(store3, TEST_KEY, 300, 3000, 999, 999, currentTimestamp - 3700); // Stale

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should only use fresh data (stores 1 and 2)
        // Average: (1 + 2) / 2 = 1, (2 + 3) / 2 = 2
        assertEq(result.numerator, 1, "Numerator should be average of 1 and 2");
        assertEq(result.denominator, 2, "Denominator should be average of 2 and 3");
    }

    function test_NumeratorDenominator_WithFailingOracle() public {
        // Test numerator/denominator when one oracle fails
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 2, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 3, 4, timestamp);
        // store3 has no data (will revert)

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should compute median from successful calls (stores 1 and 2)
        // Average: (1 + 3) / 2 = 2, (2 + 4) / 2 = 3
        assertEq(result.numerator, 2, "Numerator should be average of 1 and 3");
        assertEq(result.denominator, 3, "Denominator should be average of 2 and 4");
    }

    function test_NumeratorDenominator_IndependentFromFairValue() public {
        // Verify that sorting preserves correspondence between all value arrays
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Store values where fairValue order doesn't match numerator/denominator order
        setStoreValues(store1, TEST_KEY, 300, 3000, 1, 2, timestamp);  // High fairValue, low num/den
        setStoreValues(store2, TEST_KEY, 200, 2000, 10, 20, timestamp); // Medium fairValue, medium num/den
        setStoreValues(store3, TEST_KEY, 100, 1000, 100, 200, timestamp); // Low fairValue, high num/den

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted by fairValue: [(100,1000,100,200), (200,2000,10,20), (300,3000,1,2)]
        // Median (middle element): fairValue=200, numerator=10, denominator=20
        // Note: Arrays are sorted together, preserving correspondence
        assertEq(result.fairValue, 200, "FairValue median should be 200");
        assertEq(result.numerator, 10, "Numerator median should be 10");
        assertEq(result.denominator, 20, "Denominator median should be 20");
    }

    function test_NumeratorDenominator_GetValueIntegration() public {
        // Test that getValue correctly returns numerator and denominator
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 2, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 3, 4, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 5, 6, timestamp);

        // Test getValue with numerator key
        (uint128 numValue, uint128 numTs) = oracle.getValue("numerator:BTC/USD");
        assertEq(uint256(numValue), 3, "getValue should return median numerator");
        assertEq(uint256(numTs), block.timestamp, "Timestamp should match");

        // Test getValue with denominator key
        (uint128 denValue, uint128 denTs) = oracle.getValue("denominator:BTC/USD");
        assertEq(uint256(denValue), 4, "getValue should return median denominator");
        assertEq(uint256(denTs), block.timestamp, "Timestamp should match");
    }

    function test_NumeratorDenominator_DataDrivenOdd() public {
        OracleData[] memory oracleData = new OracleData[](3);
        oracleData[0] = OracleData(100, 1000, 5, 10, 1000000, false);
        oracleData[1] = OracleData(200, 2000, 10, 20, 1000000, false);
        oracleData[2] = OracleData(300, 3000, 15, 30, 1000000, false);

        TestCase memory testCase = TestCase({
            name: "Numerator/Denominator - Odd Oracles",
            oracleValues: oracleData,
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 200,
            expectedUsdValue: 2000,
            expectedNumerator: 10,    // Middle value
            expectedDenominator: 20,  // Middle value
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });

        runTestCase(testCase);
    }

    function test_NumeratorDenominator_DataDrivenEven() public {
        OracleData[] memory oracleData = new OracleData[](4);
        oracleData[0] = OracleData(100, 1000, 2, 4, 1000000, false);
        oracleData[1] = OracleData(200, 2000, 4, 8, 1000000, false);
        oracleData[2] = OracleData(300, 3000, 6, 12, 1000000, false);
        oracleData[3] = OracleData(400, 4000, 8, 16, 1000000, false);

        TestCase memory testCase = TestCase({
            name: "Numerator/Denominator - Even Oracles",
            oracleValues: oracleData,
            threshold: 2,
            timeoutSeconds: 3600,
            blockTimestamp: 1000000,
            expectedFairValue: 250,
            expectedUsdValue: 2500,
            expectedNumerator: 5,    // Average of 4 and 6
            expectedDenominator: 10, // Average of 8 and 12
            shouldRevert: false,
            expectedError: bytes4(0),
            expectedErrorCount: 0
        });

        runTestCase(testCase);
    }
}


contract AdminFunctionsTest is BaseTest {
    function test_ConstructorWithZeroAddress() public {
        vm.expectRevert(); // OwnableInvalidOwner
        new DIAOracleV3MetaFairValueField(address(0));
    }

    function test_ConstructorWithValidOwner() public {
        DIAOracleV3MetaFairValueField newOracle = new DIAOracleV3MetaFairValueField(owner);
        assertEq(newOracle.owner(), owner);
    }

    function test_TransferOwnership() public {
        address newOwner = address(0x3);

        vm.prank(owner);
        oracle.transferOwnership(newOwner);

        assertEq(oracle.owner(), newOwner);
    }

    function test_TransferOwnershipToZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(); // OwnableInvalidOwner
        oracle.transferOwnership(address(0));
    }

    function test_TransferOwnershipNotOwner() public {
        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        oracle.transferOwnership(address(0x3));
    }

    function test_AddValueStore() public {
        MockValueStore store = new MockValueStore();

        vm.prank(owner);
        oracle.addValueStore(address(store));

        assertEq(oracle.numValueStores(), 1);
        assertEq(oracle.valueStores(0), address(store));
    }

    function test_AddValueStoreZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(DIAOracleV3MetaFairValueField.ZeroAddress.selector));
        oracle.addValueStore(address(0));
    }

    function test_AddDuplicateValueStore() public {
        MockValueStore store = new MockValueStore();

        vm.prank(owner);
        oracle.addValueStore(address(store));

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(DIAOracleV3MetaFairValueField.OracleExists.selector));
        oracle.addValueStore(address(store));
    }

    function test_AddValueStoreNotOwner() public {
        MockValueStore store = new MockValueStore();

        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        oracle.addValueStore(address(store));
    }

    function test_RemoveValueStore() public {
        MockValueStore store = new MockValueStore();

        vm.prank(owner);
        oracle.addValueStore(address(store));

        assertEq(oracle.numValueStores(), 1);

        vm.prank(owner);
        oracle.removeValueStore(address(store));

        assertEq(oracle.numValueStores(), 0);
    }

    function test_RemoveValueStoreNotFound() public {
        MockValueStore store = new MockValueStore();

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(DIAOracleV3MetaFairValueField.OracleNotFound.selector));
        oracle.removeValueStore(address(store));
    }

    function test_RemoveValueStoreNotOwner() public {
        MockValueStore store = new MockValueStore();

        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        oracle.removeValueStore(address(store));
    }

    function test_RemoveMiddleValueStore() public {
        // Add 3 stores
        MockValueStore store1 = new MockValueStore();
        MockValueStore store2 = new MockValueStore();
        MockValueStore store3 = new MockValueStore();

        vm.prank(owner);
        oracle.addValueStore(address(store1));
        vm.prank(owner);
        oracle.addValueStore(address(store2));
        vm.prank(owner);
        oracle.addValueStore(address(store3));

        assertEq(oracle.numValueStores(), 3);

        // Remove middle store
        vm.prank(owner);
        oracle.removeValueStore(address(store2));

        assertEq(oracle.numValueStores(), 2);
        // Store1 should still be at index 0
        assertEq(oracle.valueStores(0), address(store1));
        // Store3 should now be at index 1 (moved from index 2)
        assertEq(oracle.valueStores(1), address(store3));
    }

    function test_SetThreshold() public {
        vm.prank(owner);
        oracle.setThreshold(5);

        assertEq(oracle.threshold(), 5);
    }

    function test_SetThresholdEmitsEvent() public {
        vm.prank(owner);
        oracle.setThreshold(3);

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit DIAOracleV3MetaFairValueField.ThresholdChanged(3, 5);
        oracle.setThreshold(5);
    }

    function test_SetThresholdNotOwner() public {
        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        oracle.setThreshold(5);
    }

    function test_SetTimeoutSeconds() public {
        vm.prank(owner);
        oracle.setTimeoutSeconds(7200);

        assertEq(oracle.timeoutSeconds(), 7200);
    }

    function test_SetTimeoutSecondsNotOwner() public {
        vm.prank(user);
        vm.expectRevert(); // OwnableUnauthorizedAccount
        oracle.setTimeoutSeconds(7200);
    }

    function test_SetTimeoutSecondsExceedsLimit() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(DIAOracleV3MetaFairValueField.TimeoutExceedsLimit.selector, 86401));
        oracle.setTimeoutSeconds(86401);
    }

    function test_SetTimeoutSecondsEmitsEvent() public {
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit DIAOracleV3MetaFairValueField.TimeoutSecondsChanged(3600, 7200);
        oracle.setTimeoutSeconds(7200);
    }

    function test_SetTimeoutSecondsAtBoundary() public {
        vm.prank(owner);
        oracle.setTimeoutSeconds(86400);
        assertEq(oracle.timeoutSeconds(), 86400);
    }
}

 
contract GetValueTest is BaseTest {
    function setUp() public override {
        super.setUp();
        // Initialize threshold and timeout for getValue tests
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
    }

    function test_GetValueWithFairKey() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 1, timestamp);
        setStoreValues(store3, "BTC", 300, 3000, 3, 1, timestamp);

        // Test getValue with "fairValue:BTC" key
        (uint128 value, uint128 ts) = oracle.getValue("fairValue:BTC");

        assertEq(uint256(value), 200); // Median fairValue
        assertEq(uint256(ts), block.timestamp);
    }

    function test_GetValueWithUsdKey() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 1, timestamp);
        setStoreValues(store3, "BTC", 300, 3000, 3, 1, timestamp);

        // Test getValue with "usdValue:BTC" key
        (uint128 value, uint128 ts) = oracle.getValue("usdValue:BTC");

        assertEq(uint256(value), 2000); // Median usdValue
        assertEq(uint256(ts), block.timestamp);
    }

    function test_GetValueWithNumeratorKey() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 1, timestamp);
        setStoreValues(store3, "BTC", 300, 3000, 3, 1, timestamp);

        // Test getValue with "numerator:BTC" key
        (uint128 value, uint128 ts) = oracle.getValue("numerator:BTC");

        assertEq(uint256(value), 2); // Median numerator
        assertEq(uint256(ts), block.timestamp);
    }

    function test_GetValueWithDenominatorKey() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;
        setStoreValues(store1, "BTC", 100, 1000, 1, 2, timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 3, timestamp);
        setStoreValues(store3, "BTC", 300, 3000, 3, 4, timestamp);

        // Test getValue with "denominator:BTC" key
        (uint128 value, uint128 ts) = oracle.getValue("denominator:BTC");

        assertEq(uint256(value), 3); // Median denominator
        assertEq(uint256(ts), block.timestamp);
    }

    function test_GetValueWithInvalidKey() public view {
        // Test with key that has no colon
        (uint128 value, uint128 ts) = oracle.getValue("InvalidKey");

        assertEq(uint256(value), 0);
        assertEq(uint256(ts), 0);
    }

    function test_GetValueWithUnknownAction() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, block.timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 1, block.timestamp);

        // Test getValue with unknown action - should return fairValue as default
        (uint128 value, uint128 ts) = oracle.getValue("unknownAction:BTC");

        assertEq(uint256(value), 150); // Default to fairValue (median of 100 and 200)
        assertEq(uint256(ts), block.timestamp);
    }

    function test_GetValueWithKeyStartingWithColon() public {
        // Test with key that starts with colon - no action before colon
        // First add stores with data
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, block.timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 2, 1, block.timestamp);

        // actionHash will be keccak256("") which is not bytes32(0), but it won't match any known action
        // So it defaults to fairValue
        (uint128 value, uint128 ts) = oracle.getValue(":BTC");
        assertEq(uint256(value), 150); // Median
    }

    function test_GetValueWithMultipleColons() public {
        // Setup stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        setStoreValues(store1, "BTC:USD", 100, 1000, 1, 1, block.timestamp);
        setStoreValues(store2, "BTC:USD", 200, 2000, 2, 1, block.timestamp);

        // Test with key containing multiple colons
        // Should parse "fairValue" as action and "BTC:USD" as asset
        (uint128 value, uint128 ts) = oracle.getValue("fairValue:BTC:USD");

        assertEq(uint256(value), 150); // Median of 100 and 200
        assertEq(uint256(ts), block.timestamp);
    }
}

/*//////////////////////////////////////////////////////////////
                        TIMESTAMP TEST SUITE
//////////////////////////////////////////////////////////////*/
contract TimestampTest is BaseTest {
    function setUp() public override {
        super.setUp();
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
    }

    function test_Timestamp_ReturnsOracleData_OddOracles() public {
        // Test with odd number of oracles - should return median timestamp
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 500);
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, baseTimestamp + 1000);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should return median value (200) and its timestamp (baseTimestamp + 500)
        assertEq(result.fairValue, 200);
        assertEq(result.timestamp, baseTimestamp + 500, "Should return median oracle timestamp");
    }

    function test_Timestamp_ReturnsOracleData_EvenOracles() public {
        // Test with even number of oracles - should return average of middle two timestamps
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 400);
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, baseTimestamp + 600);
        setStoreValues(store4, TEST_KEY, 400, 4000, 4, 1, baseTimestamp + 1000);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Median values: 200 and 300
        // Average timestamps: (baseTimestamp + 400 + baseTimestamp + 600) / 2 = baseTimestamp + 500
        assertEq(result.fairValue, 250);
        assertEq(result.timestamp, baseTimestamp + 500, "Should return average of middle two timestamps");
    }

    function test_Timestamp_TrackedThroughSorting() public {
        // Verify that timestamp moves with the value during sorting
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        // Values: 300 (oldest), 200 (middle), 100 (newest)
        setStoreValues(store1, TEST_KEY, 300, 3000, 3, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 500);
        setStoreValues(store3, TEST_KEY, 100, 1000, 1, 1, baseTimestamp + 1000);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted by fairValue: [100, 200, 300]
        // Sorted timestamps: [baseTimestamp + 1000, baseTimestamp + 500, baseTimestamp]
        // Median (middle): fairValue=200, timestamp=baseTimestamp + 500
        assertEq(result.fairValue, 200);
        assertEq(result.timestamp, baseTimestamp + 500, "Timestamp should follow value during sorting");
    }

    function test_Timestamp_WithStaleDataExcluded() public {
        // Verify that stale data is excluded and doesn't affect timestamp
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 currentTimestamp = 1000000;
        vm.warp(currentTimestamp);
        
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, currentTimestamp);           // Fresh
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, currentTimestamp + 500);           // Fresh
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, currentTimestamp - 3700); // Stale

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should only use fresh data (stores 1 and 2)
        // Average of timestamps: (currentTimestamp + currentTimestamp + 500) / 2 = currentTimestamp + 250
        assertEq(result.fairValue, 150);
        assertEq(result.timestamp, currentTimestamp + 250, "Should average only fresh timestamps");
    }

    function test_Timestamp_GetValueIntegration() public {
        // Test that getValue returns the median timestamp
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 500);
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, baseTimestamp + 1000);

        // Test getValue with fairValue key
        (uint128 value, uint128 ts) = oracle.getValue("fairValue:BTC/USD");
        assertEq(uint256(value), 200);
        assertEq(uint256(ts), baseTimestamp + 500, "getValue should return median timestamp");
    }

    function test_Timestamp_SameValuesDifferentTimestamps() public {
        // Test when values are the same but timestamps differ
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        // Same fairValue, different timestamps
        setStoreValues(store1, TEST_KEY, 200, 2000, 2, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 500);
        setStoreValues(store3, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 1000);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // All values are 200, so median is 200
        // Timestamps sorted: [baseTimestamp, baseTimestamp + 500, baseTimestamp + 1000]
        // Median timestamp: baseTimestamp + 500
        assertEq(result.fairValue, 200);
        assertEq(result.timestamp, baseTimestamp + 500, "Should use median timestamp even when values are same");
    }

    function test_Timestamp_LargeTimeDifferences() public {
        // Test with large differences in oracle timestamps
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 baseTimestamp = block.timestamp;
        
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, baseTimestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, baseTimestamp + 3600);  // 1 hour later
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, baseTimestamp + 7200);  // 2 hours later

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Median timestamp should be baseTimestamp + 3600
        assertEq(result.fairValue, 200);
        assertEq(result.timestamp, baseTimestamp + 3600, "Should handle large time differences");
    }
}
