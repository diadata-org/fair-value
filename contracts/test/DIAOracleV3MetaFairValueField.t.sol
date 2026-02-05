// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.20;

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

    function setUp() public {
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
            } else if (testCase.expectedError == DIAOracleV3MetaFairValueField.InvalidTimeout.selector) {
                // Need to set threshold first, then expect revert in setTimeoutSeconds
                vm.prank(owner);
                oracle.setThreshold(testCase.threshold > 0 ? testCase.threshold : 1);

                vm.prank(owner);
                vm.expectRevert(
                    abi.encodeWithSelector(
                        DIAOracleV3MetaFairValueField.InvalidTimeout.selector,
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
            expectedError: DIAOracleV3MetaFairValueField.InvalidTimeout.selector,
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
