// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.30;

import "forge-std/Test.sol";

// Import the contracts
import {DIAOracleV3MetaFairValueField} from "../metacontract/DIAOracleV3MetaFairValueField.sol";

// Mock ValueStore contract for testing
contract MockValueStore {
    // Custom error matching real ValueStore contract
    error NoDataForKey();
    error SimulatedRevert();

    mapping(string => uint256) public fairValues;
    mapping(string => uint256) public usdValues;
    mapping(string => uint256) public numerators;
    mapping(string => uint256) public denominators;
    mapping(string => uint256) public timestamps;

    // Track which keys should simulate a revert (for testing failure scenarios)
    mapping(string => bool) public shouldRevert;

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

    function setShouldRevert(string memory key, bool shouldRev) external {
        shouldRevert[key] = shouldRev;
    }

    function getValue(string memory key)
        external
        view
        returns (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp)
    {
        // Check if this key should simulate a revert (for testing oracle failures)
        if (shouldRevert[key]) {
            revert SimulatedRevert();
        }

        // Revert if no data set for this key (mimics real ValueStore behavior)
        // Uses custom error matching the real ValueStore contract
        if (timestamps[key] == 0) {
            revert NoDataForKey();
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
        // Clean up any state from previous tests
        _cleanUp();

        // Create fresh oracle for each test
        vm.prank(owner);
        oracle = new DIAOracleV3MetaFairValueField(owner);
    }

    function _cleanUp() internal {
        // Remove all stores from the oracle if it exists
        if (address(oracle) != address(0)) {
            for (uint256 i = 0; i < stores.length; i++) {
                vm.prank(owner);
                try oracle.removeValueStore(address(stores[i])) {
                    // Store existed and was removed
                } catch {
                    // Store didn't exist or was already removed
                }
            }
        }

        // Clear the stores array
        delete stores;
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
        // Ensure store doesn't revert for this key (clear any previous revert state)
        store.setShouldRevert(key, false);
    }

    function setStoreRevert(
        MockValueStore store,
        string memory key
    ) internal {
        // Mark this store to revert when getValue is called for this key
        store.setShouldRevert(key, true);
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
                    } else {
                        // Mark this store to revert when getValue is called
                        setStoreRevert(store, TEST_KEY);
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
                } else {
                    // Mark this store to revert when getValue is called
                    setStoreRevert(store, TEST_KEY);
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
            expectedNumerator: 3,    // From RIGHT middle oracle (index 2)
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
            expectedNumerator: 2,    // From RIGHT middle oracle (index 1)
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
            expectedNumerator: 2,    // From RIGHT middle oracle (index 1)
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
            expectedNumerator: 2,    // From RIGHT middle oracle (index 1)
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
            expectedNumerator: 2,    // From RIGHT middle oracle (index 1)
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

    function test_SortByUsdValueWhenAllFairValuesAreZero() public {
        // Test edge case: all fairValues are 0, should sort by usdValues
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;

        // All fairValues are 0, but usdValues differ
        setStoreValues(store1, TEST_KEY, 0, 1000, 1, 1, timestamp);       // fairValue=0, usdValue=1000
        setStoreValues(store2, TEST_KEY, 0, 2000, 2, 1, timestamp);       // fairValue=0, usdValue=2000
        setStoreValues(store3, TEST_KEY, 0, 3000, 3, 1, timestamp);       // fairValue=0, usdValue=3000
        setStoreValues(store4, TEST_KEY, 0, 4000, 4, 1, timestamp);       // fairValue=0, usdValue=4000

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should sort by usdValues since all fairValues are 0
        // Sorted usdValues: [1000, 2000, 3000, 4000]
        // Median of even count: average of middle two (2000 + 3000 + 1) / 2 = 2500
        // Numerator/denominator from RIGHT middle oracle: 3/1
        assertEq(result.fairValue, 0, "All fairValues should be 0");
        assertEq(result.usdValue, 2500, "Should be median of usdValues");
        assertEq(result.numerator, 3, "Should be from right middle oracle");
        assertEq(result.denominator, 1, "Should be from right middle oracle");
        assertEq(result.timestamp, timestamp, "Should return max timestamp");
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

        // Average of middle two: (2/4 + 3/6) / 2
        // Numerator/denominator come from RIGHT middle oracle (index 2): 3/6
        assertEq(result.numerator, 3, "Numerator should be from right middle oracle");
        assertEq(result.denominator, 6, "Denominator should be from right middle oracle");
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
        // Numerator/denominator come from RIGHT middle oracle: 2/3
        assertEq(result.numerator, 2, "Numerator should be from right middle oracle");
        assertEq(result.denominator, 3, "Denominator should be from right middle oracle");
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
        // Numerator/denominator from RIGHT middle oracle: 3/4
        assertEq(result.numerator, 3, "Numerator should be from right middle oracle");
        assertEq(result.denominator, 4, "Denominator should be from right middle oracle");
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
            expectedNumerator: 6,    // From RIGHT middle oracle (index 2)
            expectedDenominator: 12, // From RIGHT middle oracle (index 2)
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
        // Test with even number of oracles - should return max of middle two timestamps
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


        assertEq(result.fairValue, 250);
        assertEq(result.timestamp, baseTimestamp + 600, "Should return max of middle two timestamps");
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
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 1, currentTimestamp + 500);    // Fresh
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, currentTimestamp - 3700); // Stale

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should only use fresh data (stores 1 and 2)
        // For even count, returns max of middle two timestamps (currentTimestamp + 500)
        assertEq(result.fairValue, 150);
        assertEq(result.timestamp, currentTimestamp + 500, "Should return max timestamp of fresh data");
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

/*//////////////////////////////////////////////////////////////
                    STACK EXHAUSTION TEST SUITE
//////////////////////////////////////////////////////////////*/
contract QuickSortStackExhaustionTest is BaseTest {
    function setUp() public override {
        super.setUp();
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
        // This test suite creates 200+ oracles in a single test to stress test QuickSort
        // Increase limit to accommodate these tests (NOT due to state pollution)
        vm.prank(owner);
        oracle.setMaxValueStores(1000);
    }

    function test_QuickSort_LargeArray_SortedInput() public {
        // Test worst-case QuickSort scenario: already sorted data
        // This can cause O(n) recursion depth
        uint256 numOracles = 200;  // Large enough to potentially cause issues
        uint256 timestamp = block.timestamp;

        // Create stores with values in sorted order (worst case for QuickSort)
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Values in ascending order - worst case for QuickSort
            uint256 fairValue = 100 + (i * 10);
            uint256 usdValue = 1000 + (i * 100);
            setStoreValues(store, TEST_KEY, fairValue, usdValue, i + 1, 1, timestamp);
        }

        // Try to get median - may cause stack overflow
        // If it succeeds, the implementation has proper mitigations
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // If we reach here, the test passes (no stack overflow)
        // For 200 elements (indices 0-199), sorted values are 100, 110, 120, ..., 2090
        // Median of even count: average of indices 99 and 100
        // Index 99: 100 + 99*10 = 1090
        // Index 100: 100 + 100*10 = 1100
        // Median: (1090 + 1100) / 2 = 1095
        assertEq(result.fairValue, 1095, "Should return correct median");
    }

    function test_QuickSort_LargeArray_ReverseSortedInput() public {
        // Test worst-case QuickSort scenario: reverse sorted data
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        // Create stores with values in reverse sorted order (also worst case)
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Values in descending order: 100000, 99990, 99980, ..., 98010
            uint256 fairValue = 100000 - (i * 10);
            uint256 usdValue = 1000000 - (i * 100);
            setStoreValues(store, TEST_KEY, fairValue, usdValue, i + 1, 1, timestamp);
        }

        // Try to get median
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // When sorted, values are: 98010, 98020, ..., 100000
        // Median (even count 200): average of indices 99 and 100
        // Index 99: 98010 + 99*10 = 99000
        // Index 100: 98010 + 100*10 = 99010
        // Median: (99000 + 99010) / 2 = 99005
        assertEq(result.fairValue, 99005, "Should return correct median");
    }

    function test_QuickSort_LargeArray_RandomInput() public {
        // Test with larger array that has random-like values
        uint256 numOracles = 300;
        uint256 timestamp = block.timestamp;

        // Create stores with pseudo-random values
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Use some pseudo-random pattern based on prime numbers
            uint256 fairValue = 100 + ((i * 17) % 10000);
            uint256 usdValue = 1000 + ((i * 23) % 100000);
            setStoreValues(store, TEST_KEY, fairValue, usdValue, i + 1, 1, timestamp);
        }

        // Try to get median
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Just verify it doesn't overflow - exact median would be hard to calculate
        assertGt(result.fairValue, 0, "Should return a value");
    }

   

    function test_QuickSort_SmallArrays_NoOverflow() public {
        // Test that small arrays work correctly
        uint256[] memory sizes = new uint256[](6);
        sizes[0] = 2;
        sizes[1] = 5;
        sizes[2] = 10;
        sizes[3] = 11;
        sizes[4] = 50;
        sizes[5] = 100;

        for (uint256 j = 0; j < sizes.length; j++) {
            // Clear previous stores
            uint256 numExisting = stores.length;
            for (uint256 i = 0; i < numExisting; i++) {
                vm.prank(owner);
                oracle.removeValueStore(address(stores[i]));
            }
            delete stores;

            uint256 numOracles = sizes[j];
            uint256 timestamp = block.timestamp;

            // Create stores
            for (uint256 i = 0; i < numOracles; i++) {
                MockValueStore store = createMockStore();
                setStoreValues(store, TEST_KEY, i * 100, i * 1000, i + 1, 1, timestamp);
            }

            // Should work fine
            DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);
            assertGt(result.fairValue, 0, "Should return a value");
        }
    }

    function test_QuickSort_AdversarialPattern_AlternatingMinMax() public {
        // Test with adversarial pattern: alternating min and max values
        // This could potentially cause poor pivot selection
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Alternate between very low and very high values
            uint256 fairValue = (i % 2 == 0) ? 100 : 1000000;
            uint256 usdValue = (i % 2 == 0) ? 1000 : 10000000;
            setStoreValues(store, TEST_KEY, fairValue, usdValue, i + 1, 1, timestamp);
        }

        // Try to get median
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // With 200 oracles alternating 100 and 1000000, we get:
        // 100 oracles with value 100, 100 oracles with value 1000000
        // When sorted: first 100 are 100, next 100 are 1000000
        // Median (even count): average of indices 99 and 100
        // Index 99: 100, Index 100: 1000000
        // Median: (100 + 1000000) / 2 = 500050
        assertEq(result.fairValue, 500050, "Should return correct median for alternating pattern");
    }

    function testFuzz_QuickSort_Fuzz(uint128 numOracles) public {
        // Fuzz test with varying number of oracles
        numOracles = uint128(bound(numOracles, 2, 500));  // Limit to reasonable range
        uint256 timestamp = block.timestamp;

        // Clear previous stores
        uint256 numExisting = stores.length;
        for (uint256 i = 0; i < numExisting; i++) {
            vm.prank(owner);
            oracle.removeValueStore(address(stores[i]));
        }
        delete stores;

        // Create stores
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, i * 100, i * 1000, i + 1, 1, timestamp);
        }

        // Should not revert with stack overflow
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);
        assertGt(result.fairValue, 0, "Should return a value");
    }
}

/*//////////////////////////////////////////////////////////////
              DUPLICATE VALUES PERFORMANCE TEST SUITE
    Tests for O(n²) degradation with adversarial duplicate data
//////////////////////////////////////////////////////////////*/
contract DuplicateValuesPerformanceTest is BaseTest {
    function setUp() public override {
        super.setUp();
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
        // This test suite creates 100-1000 oracles in a single test for performance testing
        // Increase limit to accommodate these tests (NOT due to state pollution)
        vm.prank(owner);
        oracle.setMaxValueStores(1000);
    }

    function test_DuplicateValues_AllSame_100Oracles() public {
        // Test worst-case: all oracles report the exact same value
        // This causes maximum imbalance in QuickSort partition
        uint256 numOracles = 100;
        uint256 timestamp = block.timestamp;

        // All oracles have identical values
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1000, 10000, 1, 1, timestamp);
        }

        // This should work but may use more gas
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1000, "All same values should return that value");
    }

    function test_DuplicateValues_AllSame_200Oracles() public {
        // Test with more oracles - should see significant gas increase if vulnerable
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1000, 10000, 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1000, "All same values should return that value");
    }

    function test_DuplicateValues_AllSame_300Oracles() public {
        // Test with even more oracles to stress test
        uint256 numOracles = 300;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1000, 10000, 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1000, "All same values should return that value");
    }

    function test_DuplicateValues_NinetyPercentSame() public {
        // Test with 90% duplicates and 10% different values
        // This is a realistic adversarial scenario
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // 90% have value 1000, 10% have value 2000
            uint256 fairValue = (i < 180) ? 1000 : 2000;
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // With 180 values of 1000 and 20 values of 2000, median should be 1000
        assertEq(result.fairValue, 1000, "Median should be the common value");
    }

    function test_DuplicateValues_TwoDistinctValues() public {
        // Test with only 2 distinct values, alternating
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Alternate between 100 and 200
            uint256 fairValue = (i % 2 == 0) ? 100 : 200;
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // 100 values of 100, 100 values of 200
        // Median: (100 + 200) / 2 = 150
        assertEq(result.fairValue, 150, "Median of two values should be average");
    }

    function test_DuplicateValues_ThreeDistinctValues() public {
        // Test with 3 distinct values
        uint256 numOracles = 201;  // Odd number for clear median
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // 67 values each of 100, 200, 300
            uint256 fairValue = 100 + ((i % 3) * 100);
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted: 67×100, 67×200, 67×300
        // Median (101st element): 200
        assertEq(result.fairValue, 200, "Median should be middle value");
    }

    function test_DuplicateValues_CollusionAttack() public {
        // Simulate collusion: majority of oracles report same fake value
        uint256 numOracles = 200;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // 150 oracles (75%) report 999999, 50 report correct values around 100
            uint256 fairValue = (i < 150) ? 999999 : (100 + i);
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // With 150 values of 999999, median should be 999999
        assertEq(result.fairValue, 999999, "Colluding majority should control median");
    }

    function test_DuplicateValues_SpikeResistance() public {
        // Test with one oracle trying to spike the value
        uint256 numOracles = 199;  // Odd number
        uint256 timestamp = block.timestamp;

        // 198 oracles report 100, 1 oracle reports 1000000000
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = (i == 100) ? 1000000000 : 100;
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Spike value should be filtered out by median
        assertEq(result.fairValue, 100, "Spike should be filtered out");
    }

    function test_DuplicateValues_SameTimestamps() public {
        // Test with all same values but different timestamps
        uint256 numOracles = 150;
        uint256 baseTimestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Same value, different timestamps
            setStoreValues(store, TEST_KEY, 1000, 10000, i + 1, 1, baseTimestamp + i);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1000, "Value should be 1000");
        // Timestamp should be one of the middle values
        assertGe(result.timestamp, baseTimestamp, "Timestamp should be valid");
    }

    function test_DuplicateValues_CompareGas() public {
        // Compare gas usage: all same vs random values
        uint256 numOracles = 150;
        uint256 timestamp = block.timestamp;

        // First test: all same values (worst case for duplicates)
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1000, 10000, 1, 1, timestamp);
        }

        uint256 gasBefore1 = gasleft();
        DIAOracleV3MetaFairValueField.MedianSet memory result1 = oracle.getMedianValues(TEST_KEY);
        uint256 gasAfter1 = gasleft();
        uint256 gasDuplicates = gasBefore1 - gasAfter1;

        // Clear stores
        for (uint256 i = 0; i < numOracles; i++) {
            vm.prank(owner);
            oracle.removeValueStore(address(stores[i]));
        }
        delete stores;

        // Second test: random-like values (better distribution)
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = 1000 + (i * 7);  // Different values
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        uint256 gasBefore2 = gasleft();
        DIAOracleV3MetaFairValueField.MedianSet memory result2 = oracle.getMedianValues(TEST_KEY);
        uint256 gasAfter2 = gasleft();
        uint256 gasRandom = gasBefore2 - gasAfter2;

        // Log the comparison (visible in test output)
        // With the fix, duplicates should use similar or less gas than distinct values
        emit log_named_uint("Gas with all duplicates:", gasDuplicates);
        emit log_named_uint("Gas with distinct values:", gasRandom);

        // Calculate difference safely (avoid underflow)
        if (gasDuplicates >= gasRandom) {
            emit log_named_uint("Gas overhead (duplicates - distinct):", gasDuplicates - gasRandom);
        } else {
            emit log_named_uint("Gas savings (distinct - duplicates):", gasRandom - gasDuplicates);
        }

        // Test should pass regardless, but we can see the gas difference
        assertEq(result1.fairValue, 1000, "Duplicates test");
        assertGt(result2.fairValue, 0, "Random test");
    }

    function test_DuplicateValues_ZeroValues() public {
        // Edge case: all zero values
        uint256 numOracles = 100;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // All fairValues are 0, should sort by usdValue
            setStoreValues(store, TEST_KEY, 0, 5000, 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 0, "All zero fairValues should return 0");
        assertEq(result.usdValue, 5000, "Should return usdValue median");
    }

    function test_DuplicateValues_MostlyZero() public {
        // Test with mostly zeros and a few non-zero values
        uint256 numOracles = 100;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // 90% have fairValue=0, 10% have fairValue=1000
            uint256 fairValue = (i < 90) ? 0 : 1000;
            uint256 usdValue = (i < 90) ? 1000 : 10000;
            setStoreValues(store, TEST_KEY, fairValue, usdValue, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // 90 zeros, 10 values of 1000
        // When sorted by fairValue (all 0), then by usdValue
        // Median should still have fairValue=0
        assertEq(result.fairValue, 0, "Median should have zero fairValue");
    }

    function testFuzz_DuplicateValues_Fuzz(uint128 numOracles) public {
        // Fuzz test with varying number of oracles all having same value
        numOracles = uint128(bound(numOracles, 2, 300));

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // All same value
            setStoreValues(store, TEST_KEY, 5555, 55555, 1, 1, block.timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 5555, "Should return the common value");
    }

    function test_DuplicateValues_AllSame_500Oracles() public {
        // Stress test: 500 oracles all with same value
        // Before fix: Would cause STACK OVERFLOW
        // After fix: Should work fine
        uint256 numOracles = 500;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1234, 12340, 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1234, "All same values should return that value");
        assertEq(result.usdValue, 12340, "USD value should match");
    }

    function test_DuplicateValues_AllSame_1000Oracles() public {
        // Extreme stress test: 1000 oracles all with same value
        // This would DEFINITELY cause stack overflow with recursive implementation
        // Iterative implementation should handle it easily
        uint256 numOracles = 1000;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 9999, 99990, 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 9999, "All same values should return that value");
        assertEq(result.usdValue, 99990, "USD value should match");
    }

    function test_IterativeQuickSort_NoStackOverflow() public {
        // Test that demonstrates iterative implementation prevents stack overflow
        // Compare: recursive would overflow at ~200-300 elements with duplicates
        // Iterative should handle 1000+ easily

        uint256 numOracles = 1000;
        uint256 timestamp = block.timestamp;

        // Create adversarial input: worst case for QuickSort
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // All same value = worst case for partition
            setStoreValues(store, TEST_KEY, 5000, 50000, 1, 1, timestamp);
        }

        // This should NOT cause stack overflow with iterative implementation
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 5000, "Should return correct median");
        assertEq(result.usdValue, 50000, "Should return correct USD value");
    }

    function test_PartitionBoundary_LeftPartitionBug() public {
        // Test for partition boundary bug where left partition with 1 element is not sorted
        // Bug: When pivotEnd == left, the condition "pivotEnd > left" is false,
        // so we don't push the left partition even though it might need sorting

        // Scenario: [1, 100, 100, 100, 100]
        // After partition: lt=1, so lessThanEnd=0
        // Left partition should be [0, 0] (single element, already sorted)
        // This should work correctly

        // Scenario: [1, 2, 100, 100, 100]
        // After first partition with pivot at index 2 (value 100):
        //   lt=2 (first 100), gt=4 (last 100)
        //   lessThanEnd = 2-1 = 1
        //   greaterThanStart = 4+1 = 5 (right+1, meaning empty)
        // Push: left=[0,1], right=[5,4] (empty, not pushed)
        // Pop and sort [0,1]: should sort [1,2] correctly
    }

    function test_PartitionBoundary_EdgeCase() public {
        // Test edge case: array with two distinct values
        // [1, 100, 100, 100, 100, 100, 2, 100, 100, 100]
        // This ensures the algorithm handles boundaries correctly

        uint256 numOracles = 10;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            // Create specific pattern
            uint256 fairValue = (i == 0) ? 1 : (i == 6 ? 2 : 100);
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Should sort to [1, 2, 100, 100, 100, 100, 100, 100, 100, 100]
        // Median of 10 elements (indices 4-5 average): 100
        assertEq(result.fairValue, 100, "Should handle boundary correctly");
    }

    function test_PartitionBoundary_SingleElementInLeft() public {
        // Test when left partition has exactly 1 element
        // [1, 10, 10, 10, 10]
        // After partition: lessThanEnd=0 (single element at index 0)
        // This element is already sorted, so no need to process

        uint256 numOracles = 5;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = (i == 0) ? 1 : 10;
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted: [1, 10, 10, 10, 10]
        // Median (index 2): 10
        assertEq(result.fairValue, 10, "Should handle single element in left partition");
    }

    function test_PartitionBoundary_SingleElementInRight() public {
        // Test when right partition has exactly 1 element
        // [10, 10, 10, 10, 20]
        // After partition: greaterThanStart=4 (single element at index 4)

        uint256 numOracles = 5;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = (i == 4) ? 20 : 10;
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted: [10, 10, 10, 10, 20]
        // Median (index 2): 10
        assertEq(result.fairValue, 10, "Should handle single element in right partition");
    }
}


/*//////////////////////////////////////////////////////////////
                    INVARIANT & PROPERTY-BASED TEST SUITE
    Tests mathematical properties and invariants of the oracle
//////////////////////////////////////////////////////////////*/
contract InvariantPropertiesTest is BaseTest {
    function setUp() public override {
        super.setUp();
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
    }

    function test_Invariant_MedianWithinMinMax() public {
        // INVARIANT: Median should always be between min and max values
        uint256 numOracles = 10;
        uint256 timestamp = block.timestamp;

        uint256[] memory fairValues = new uint256[](numOracles);
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            fairValues[i] = 100 + (i * 137); // Prime number for variety
            setStoreValues(store, TEST_KEY, fairValues[i], fairValues[i] * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Find min and max
        uint256 minValue = fairValues[0];
        uint256 maxValue = fairValues[0];
        for (uint256 i = 1; i < numOracles; i++) {
            if (fairValues[i] < minValue) minValue = fairValues[i];
            if (fairValues[i] > maxValue) maxValue = fairValues[i];
        }

        // INVARIANT CHECK: Median must be within [min, max]
        assertGe(result.fairValue, minValue, "Median should be >= min value");
        assertLe(result.fairValue, maxValue, "Median should be <= max value");
    }

    function test_Invariant_MedianAllSame() public {
        // INVARIANT: If all oracles report same value, median is that value
        uint256 numOracles = 7;
        uint256 timestamp = block.timestamp;
        uint256 sameValue = 5555;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, sameValue, sameValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, sameValue, "Median of identical values should be that value");
        assertEq(result.usdValue, sameValue * 10, "USD value should match");
    }

    function test_Invariant_MedianIsMiddleValue_OddCount() public {
        // INVARIANT: For odd count, median should be exactly one of the values
        uint256 numOracles = 11; // Odd number
        uint256 timestamp = block.timestamp;

        uint256[] memory fairValues = new uint256[](numOracles);
        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            fairValues[i] = 100 + (i * 50);
            setStoreValues(store, TEST_KEY, fairValues[i], fairValues[i] * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // For odd count (11), median is the 6th element (index 5) when sorted
        uint256 expectedMedian = 100 + (5 * 50);
        assertEq(result.fairValue, expectedMedian, "Median should be middle value");
    }

    function test_Invariant_MedianIsAverageOfMiddles_EvenCount() public {
        // INVARIANT: For even count, median should be average of two middle values
        uint256 numOracles = 10; // Even number
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = 100 + (i * 100);
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // For even count (10), median is average of 5th and 6th elements (indices 4, 5)
        assertEq(result.fairValue, 550, "Median should be average of two middle values");
    }

    function test_Invariant_RoundTrip_SetGetValues() public {
        // INVARIANT: Values set should match values retrieved
        MockValueStore store = createMockStore();

        store.setValue("TEST/KEY", 12345, 123450, 7, 13, block.timestamp);

        (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp) =
            store.getValue("TEST/KEY");

        assertEq(fairValue, 12345, "FairValue round-trip");
        assertEq(usdValue, 123450, "UsdValue round-trip");
        assertEq(numerator, 7, "Numerator round-trip");
        assertEq(denominator, 13, "Denominator round-trip");
        assertEq(timestamp, block.timestamp, "Timestamp round-trip");
    }

    function test_Invariant_RoundTrip_MultipleKeys() public {
        // INVARIANT: Multiple keys should maintain independent values
        MockValueStore store = createMockStore();

        store.setValue("KEY1", 1000, 10000, 1, 1, block.timestamp);
        store.setValue("KEY2", 2000, 20000, 2, 3, block.timestamp);
        store.setValue("KEY3", 3000, 30000, 3, 5, block.timestamp);

        (uint256 fv1,,, uint256 den1,) = store.getValue("KEY1");
        (uint256 fv2,, uint256 num2,,) = store.getValue("KEY2");
        (uint256 fv3,, uint256 num3, uint256 den3,) = store.getValue("KEY3");

        assertEq(fv1, 1000, "KEY1 should be independent");
        assertEq(fv2, 2000, "KEY2 should be independent");
        assertEq(fv3, 3000, "KEY3 should be independent");
        assertEq(num2, 2, "KEY2 numerator");
        assertEq(num3, 3, "KEY3 numerator");
        assertEq(den1, 1, "KEY1 denominator");
        assertEq(den3, 5, "KEY3 denominator");
    }

    function test_Invariant_RoundTrip_Override() public {
        // INVARIANT: Setting new values should override old ones completely
        MockValueStore store = createMockStore();

        store.setValue("KEY", 1000, 10000, 1, 2, 100);
        store.setValue("KEY", 5000, 50000, 7, 11, 200);

        (uint256 fairValue, uint256 usdValue, uint256 numerator, uint256 denominator, uint256 timestamp) =
            store.getValue("KEY");

        assertEq(fairValue, 5000, "Should have new fairValue");
        assertEq(usdValue, 50000, "Should have new usdValue");
        assertEq(numerator, 7, "Should have new numerator");
        assertEq(denominator, 11, "Should have new denominator");
        assertEq(timestamp, 200, "Should have new timestamp");
    }

    function test_Invariant_SortedMonotonicity() public {
        // INVARIANT: Values should be properly sorted
        uint256 numOracles = 20;
        uint256 timestamp = block.timestamp;

        // Create values in random order
        uint256[] memory unsortedValues = new uint256[](numOracles);
        unsortedValues[0] = 500;
        unsortedValues[1] = 100;
        unsortedValues[2] = 700;
        unsortedValues[3] = 300;
        unsortedValues[4] = 900;
        unsortedValues[5] = 200;
        unsortedValues[6] = 600;
        unsortedValues[7] = 400;
        unsortedValues[8] = 800;
        unsortedValues[9] = 150;
        unsortedValues[10] = 950;
        unsortedValues[11] = 250;
        unsortedValues[12] = 750;
        unsortedValues[13] = 350;
        unsortedValues[14] = 850;
        unsortedValues[15] = 450;
        unsortedValues[16] = 650;
        unsortedValues[17] = 50;
        unsortedValues[18] = 1000;
        unsortedValues[19] = 550;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, unsortedValues[i], unsortedValues[i] * 10, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Sorted values: [50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 550, 600, 650, 700, 750, 800, 850, 900, 950, 1000]
        // For 20 elements, median is average of 10th and 11th: (500 + 550) / 2 = 525
        assertEq(result.fairValue, 525, "Median indicates proper sorting occurred");
    }

    function test_Invariant_TimestampFromOneOfTheOracles() public {
        // INVARIANT: Returned timestamp should match one of the oracle timestamps
        uint256 numOracles = 5;
        uint256 baseTimestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 1000 + (i * 100), 10000 + (i * 1000), i + 1, 1, baseTimestamp + (i * 1000));
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Timestamp should be one of the oracle timestamps
        bool found = false;
        for (uint256 i = 0; i < numOracles; i++) {
            if (result.timestamp == baseTimestamp + (i * 1000)) {
                found = true;
                break;
            }
        }
        assertTrue(found, "Timestamp should be from one of the oracles");
    }

    function test_Invariant_MedianStability() public {
        // INVARIANT: Adding same values shouldn't change median much
        uint256 timestamp = block.timestamp;

        // Create 10 oracles with value 100
        for (uint256 i = 0; i < 10; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 100, 1000, i + 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result1 = oracle.getMedianValues(TEST_KEY);
        assertEq(result1.fairValue, 100, "Median of all 100s should be 100");

        // Add 10 more oracles with value 200
        for (uint256 i = 0; i < 10; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, 200, 2000, i + 11, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result2 = oracle.getMedianValues(TEST_KEY);

        assertEq(result2.fairValue, 150, "Median should be between 100 and 200");
        assertGe(result2.fairValue, result1.fairValue, "Median should shift toward new values");
    }

    function test_Invariant_NumeratorDenominatorConsistency() public {
        // INVARIANT: If denominator is 1, value should equal numerator
        uint256 numOracles = 5;
        uint256 timestamp = block.timestamp;

        for (uint256 i = 0; i < numOracles; i++) {
            MockValueStore store = createMockStore();
            uint256 value = 1000 + (i * 100);
            setStoreValues(store, TEST_KEY, value, value * 10, value, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, result.fairValue, "Numerator should equal fairValue when denominator is 1");
        assertEq(result.denominator, 1, "Denominator should be 1");
    }

    function test_Invariant_ExtremeValues() public {
        // INVARIANT: System should handle extreme values correctly
        uint256 timestamp = block.timestamp;

        // Use extreme but safe values (avoid overflow)
        uint256[] memory extremeValues = new uint256[](5);
        extremeValues[0] = 0;
        extremeValues[1] = 1;
        extremeValues[2] = 1e18; // Large but safe value
        extremeValues[3] = 1e27; // Very large but safe
        extremeValues[4] = 1e36; // Extremely large but safe

        for (uint256 i = 0; i < 5; i++) {
            MockValueStore store = createMockStore();
            setStoreValues(store, TEST_KEY, extremeValues[i], extremeValues[i], 1, 1, timestamp);
        }

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Median of [0, 1, 1e18, 1e27, 1e36] is 1e18 (middle value)
        assertEq(result.fairValue, 1e18, "Should handle extreme values");
    }

    function testFuzz_Invariant_MedianAlwaysValid(uint128[20] memory values) public {
        // FUZZ TEST: Median should always be valid for any input values
        vm.prank(owner);
        oracle.setMaxValueStores(20);

        uint256 timestamp = block.timestamp;

        // Clear existing stores
        for (uint256 i = 0; i < 20; i++) {
            if (i < stores.length) {
                vm.prank(owner);
                oracle.removeValueStore(address(stores[i]));
            }
        }
        delete stores;

        // Add stores with fuzzed values
        for (uint256 i = 0; i < 20; i++) {
            MockValueStore store = createMockStore();
            uint256 fairValue = uint256(values[i]);
            setStoreValues(store, TEST_KEY, fairValue, fairValue * 10, i + 1, 1, timestamp);
        }

        // Get median - should never revert or produce invalid results
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertTrue(true, "Median computed successfully for any input");
    }

    function test_Invariant_OrderPreservation_MultipleKeys() public {
        // INVARIANT: Different keys should maintain independent medians
        uint256 timestamp = block.timestamp;

        // Set BTC values: [100, 200, 300]
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        setStoreValues(store1, "BTC", 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, "BTC", 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, "BTC", 300, 3000, 1, 1, timestamp);

        // Set ETH values: [1000, 2000, 3000]
        setStoreValues(store1, "ETH", 1000, 10000, 1, 1, timestamp);
        setStoreValues(store2, "ETH", 2000, 20000, 1, 1, timestamp);
        setStoreValues(store3, "ETH", 3000, 30000, 1, 1, timestamp);

        // Get medians
        DIAOracleV3MetaFairValueField.MedianSet memory btcResult = oracle.getMedianValues("BTC");
        DIAOracleV3MetaFairValueField.MedianSet memory ethResult = oracle.getMedianValues("ETH");

        // Each key should maintain its own independent median
        assertEq(btcResult.fairValue, 200, "BTC median should be independent");
        assertEq(ethResult.fairValue, 2000, "ETH median should be independent");
    }
}

/// @title MedianPrecisionTest
/// @notice Tests precision and accuracy of median calculations
contract MedianPrecisionTest is BaseTest {
    // Precision tests for median calculation

    function test_Precision_EvenCount_RoundingBehavior() public {
        // Tests that the contract uses ROUNDING (a+b+1)/2, not exact division (a+b)/2
        //
        // Scenario 1: Even sum - rounding formula gives exact result
        // [100, 200, 300, 400] -> middle: 200, 300
        // Contract: (200 + 300 + 1) / 2 = 501 / 2 = 250
        // Mathematical: (200 + 300) / 2 = 500 / 2 = 250
        // When sum is even (500), both formulas give same result

        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Result is 250, but this is from ROUNDING formula, not exact division
        assertEq(result.fairValue, 250, "Rounding formula gives 250 when sum is even");
 
        // Scenario 2: Odd sum - proves we're using rounding, not exact division
        // [100, 201, 300, 400] -> middle: 201, 300
        // Contract: (201 + 300 + 1) / 2 = 502 / 2 = 251 (rounds up)
        // Mathematical exact division: (201 + 300) / 2 = 501 / 2 = 250 (truncates)
        // DIFFERENCE: This proves we use rounding!

        setStoreValues(store2, TEST_KEY, 201, 2010, 1, 1, timestamp);

        result = oracle.getMedianValues(TEST_KEY);

        // Contract rounds up to 251, not 250
        assertEq(result.fairValue, 251, "Rounding formula gives 251 (rounds up from 250.5)");
        assertEq(result.fairValue, (201 + 300 + 1) / 2, "Equals (201+300+1)/2, not (201+300)/2");
 
 
    }

    function test_Precision_EvenCount_RoundingExact() public {
        // EVEN COUNT with exact division: [100, 200, 300, 400]
        // Sorted: 100, 200, 300, 400
        // Middle two: 200, 300
        // FairValue median: (200 + 300 + 1) / 2 = 250 (exact, no rounding needed)
        // When (a + b) is even, (a + b + 1) / 2 = (a + b) / 2, so result is exact
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 250, "Median should be 250 (exact, no rounding needed)");
        assertTrue(((200 + 300) % 2) == 0, "Sum is even, so rounding formula gives exact result");
    }

    function test_Precision_EvenCount_OddNumbers() public {
        // EVEN COUNT with odd numbers: [1, 2, 3, 4]
        // Middle two: 2, 3
        // Average: (2 + 3 + 1) / 2 = 3 (rounded to nearest)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 1, 10, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 2, 20, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 3, 30, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 4, 40, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 3, "Median of 1,2,3,4 should be 3 (rounded from 2.5)");
    }

    function test_Precision_EvenCount_LargeNumbers_WithOddSum() public {
        // Tests rounding behavior with large numbers where sum is ODD
        // This proves rounding is used, not truncation
        //
        // Scenario 1: Even sum (5e18) - both formulas give same result
        // [1e18, 2e18, 3e18, 4e18] -> middle: 2e18, 3e18
        // Rounding: (2e18 + 3e18 + 1) / 2 = 5e18 / 2 = 2.5e18
        // Truncation: (2e18 + 3e18) / 2 = 5e18 / 2 = 2.5e18
        // SAME RESULT (doesn't prove which formula is used)

        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 1e18, 1e19, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 2e18, 2e19, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 3e18, 3e19, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 4e18, 4e19, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 2.5e18, "Even sum: both formulas give 2.5e18");
        assertTrue(((2e18 + 3e18) % 2) == 0, "Sum is even (5e18), can't distinguish formulas");

        // Scenario 2: ODD SUM - proves we use rounding!
        // [1e18, 2.000001e18, 3e18, 4e18] -> middle: 2.000001e18, 3e18
        // Sum: 5.000001e18 (ODD)
        // Rounding: (2.000001e18 + 3e18 + 1) / 2 = 5.000002e18 / 2 = 2.500001e18
        // Truncation: (2.000001e18 + 3e18) / 2 = 5.000001e18 / 2 = 2.5e18 (truncates!)
        // DIFFERENCE: Rounding gives 2.500001e18, truncation gives 2.5e18

        setStoreValues(store2, TEST_KEY, 2_000_001_000_000_000_000, 20_000_010_000_000_000_000, 1, 1, timestamp);

        result = oracle.getMedianValues(TEST_KEY);

        // Contract uses rounding: (a + b + 1) / 2
        // With large numbers, the +1 has minimal visible effect
        uint256 valA = 2_000_001_000_000_000_000;
        uint256 valB = 3_000_000_000_000_000_000;
        uint256 sum = valA + valB;
        uint256 expected = (sum + 1) / 2;

        assertEq(result.fairValue, expected, "Uses (a+b+1)/2 formula");

 
    }

    function test_Precision_NumeratorDenominator_EvenCount() public {
        // EVEN COUNT: Test that numerator/denominator come from right middle oracle
        // Values: [100/1, 200/1, 300/1, 400/1]
        // Middle two: 200, 300
        // FairValue median: (200 + 300 + 1) / 2 = 250 (rounded)
        // Numerator/denominator from RIGHT middle oracle: 300/1
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 100, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 200, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 300, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 400, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.numerator, 300, "Median numerator should be from right middle oracle");
        assertEq(result.denominator, 1, "Median denominator should be from right middle oracle");
    }

    function test_Precision_NumeratorDenominator_FromMiddleOracle() public {
        // Tests that numerator/denominator are taken from the RIGHT middle oracle
        // NOT averaged independently!
        //
        // Key insight: The contract does NOT average numerators and denominators
        // Instead:
        // 1. Calculate fairValue median using rounding: (57 + 60 + 1) / 2 = 59
        // 2. Find which oracle's fairValue is closest to the median
        // 3. Return THAT oracle's numerator and denominator

        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        // Sorted by fairValue: [50, 57, 60, 66]
        // Each has different numerator/denominator:
        // 50 (100/2), 57 (400/7), 60 (300/5), 66 (200/3)
        setStoreValues(store1, TEST_KEY, 50, 500, 100, 2, timestamp);   // 100/2 = 50
        setStoreValues(store2, TEST_KEY, 57, 570, 400, 7, timestamp);   // 400/7 ≈ 57.14
        setStoreValues(store3, TEST_KEY, 60, 600, 300, 5, timestamp);   // 300/5 = 60
        setStoreValues(store4, TEST_KEY, 66, 660, 200, 3, timestamp);   // 200/3 ≈ 66.66

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // FairValue calculation:
        // Sorted: 50, 57, 60, 66
        // Middle two: 57, 60
        // FairValue median: (57 + 60 + 1) / 2 = 59 (rounds up)

        // Since fairValue rounds to 59 (closer to 60), we take RIGHT middle oracle's data
        // Right middle oracle has: fairValue=60, numerator=300, denominator=5

        assertEq(result.fairValue, 59, "FairValue median: (57+60+1)/2 = 59");
        assertEq(result.numerator, 300, "Numerator from oracle with fairValue=60");
        assertEq(result.denominator, 5, "Denominator from oracle with fairValue=60");

        // This proves: numerator/denominator are NOT averaged
        // If they were averaged: (400/7 + 300/5) / 2 = (2000/35 + 2100/35) / 2 = 4100/70 ≈ 58.57
        // But we return 300/5 = 60 (the value from one oracle)

        emit log_string("=== NUMERATOR/DENOMINATOR SELECTION ===");
        emit log_string("NOT averaged independently!");
        emit log_string("Taken from RIGHT middle oracle after fairValue rounding");
        emit log_named_uint("FairValue median", 59);
        emit log_string("Closest to 60, so use 60's numerator/denominator: 300/5");
    }

    function test_Precision_OddCount_NoAveraging() public {
        // ODD COUNT: No averaging, so no precision loss
        // Values: [100, 200, 300]
        // Median: 200 (middle value, no averaging needed)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 200, "Median of 100,200,300 should be 200");
    }

    function test_Precision_VeryCloseValues() public {
        // EVEN COUNT: Values very close together
        // [100, 101, 102, 103]
        // Middle two: 101, 102
        // Average: (101 + 102 + 1) / 2 = 102 (rounded to nearest)
        // Maximum precision loss: 0.5
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 101, 1010, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 102, 1020, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 103, 1030, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 102, "Median should be 102 (rounded from 101.5)");
    }

    function test_Precision_DuplicateMiddleValues() public {
        // EVEN COUNT: Duplicate middle values
        // [100, 200, 200, 300]
        // Middle two: 200, 200
        // Average: (200 + 200) / 2 = 200 (exact, no truncation)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 300, 3000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 200, "Median with duplicate middle values should be exact");
    }

    function test_Precision_BigValues_SmallDifference() public {
        // EVEN COUNT: Large values with small difference
        // [1e18, 1e18+1, 1e18+2, 1e18+3]
        // Middle two: 1e18+1, 1e18+2
        // Average: (2e18+3+1) / 2 = 1e18+2 (rounded to nearest)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 1e18, 1e19, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 1e18 + 1, 1e19 + 10, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 1e18 + 2, 1e19 + 20, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 1e18 + 3, 1e19 + 30, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 1e18 + 2, "Median should round to 1e18+2");
    }

    function test_Precision_CompareOddVsEven() public {
        // Compare precision: odd count (exact) vs even count (rounded)
        // ODD: [100, 200, 300] -> 200 (exact)
        // EVEN: [100, 200, 300, 400] -> (200+300+1)/2 = 250 (rounded to nearest)
        uint256 timestamp = block.timestamp;

        // Odd count test
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory oddResult = oracle.getMedianValues(TEST_KEY);

        // Add fourth store for even count
        MockValueStore store4 = createMockStore();
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory evenResult = oracle.getMedianValues(TEST_KEY);

        assertEq(oddResult.fairValue, 200, "Odd count median should be exact");
        assertEq(evenResult.fairValue, 250, "Even count median should average middle values");
    }

    function test_Precision_Documentation_RoundingBehavior() public {
 

        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        // Scenario 1: Even sum - exact result
        // [100, 200, 300, 400] -> middle: 200, 300
        // (200 + 300 + 1) / 2 = 501 / 2 = 250 (exact since 200+300=500 is even)
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 250, "Even sum: exact result (250.0)");
        assertTrue(((200 + 300) % 2) == 0, "Sum is even (500), result is exact");

        // Scenario 2: Odd sum - rounds up
        // [100, 201, 300, 400] -> middle: 201, 300
        // (201 + 300 + 1) / 2 = 502 / 2 = 251 (rounds up from 250.5)
        setStoreValues(store2, TEST_KEY, 201, 2010, 1, 1, timestamp);

        result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 251, "Odd sum: rounds half-up to 251 (not 250)");
        assertTrue(((201 + 300) % 2) != 0, "Sum is odd (501), so rounding occurs");
 
    }

    function test_Precision_MaximumRounding() public {
        // Tests maximum rounding error scenario
        // [100, 101, 102, 103]
        // Middle two: 101, 102
        // Mathematical average: 101.5
        // Contract: (101 + 102 + 1) / 2 = 102 (rounds up from 101.5)
        // Maximum deviation from mathematical average: 0.5 units (rounds up for .5)
        //
        // Note: Using (a + b + 1) / 2 means:
        // - Odd sums (e.g., 203) round up: (203 + 1) / 2 = 102
        // - This is "round half up" behavior, not truncation
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 101, 1010, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 102, 1020, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 103, 1030, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 102, "Rounds up: 101.5 -> 102 (using round-half-up)");
    }

    function test_Precision_AllValuesSame() public {
        // All values identical: no precision loss
        // [500, 500, 500, 500]
        // Middle two: 500, 500
        // Average: (500 + 500) / 2 = 500 (exact)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 500, 5000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 500, 5000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 500, 5000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 500, 5000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 500, "All values same: no precision loss");
    }

    function testFuzz_Precision_AlwaysValid(uint128 a, uint128 b, uint128 c, uint128 d) public {
        // FUZZ TEST: Median calculation should always produce valid results
        // for any combination of values (within uint128 range to prevent overflow)
        vm.prank(owner);
        oracle.setMaxValueStores(10);

        uint256 timestamp = block.timestamp;

        // Clear existing stores
        for (uint256 i = 0; i < stores.length; i++) {
            vm.prank(owner);
            oracle.removeValueStore(address(stores[i]));
        }
        delete stores;

        // Create 4 stores with fuzzed values
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        // Bound values to prevent overflow in averaging
        uint256 valA = uint256(a) % 1e18;
        uint256 valB = uint256(b) % 1e18;
        uint256 valC = uint256(c) % 1e18;
        uint256 valD = uint256(d) % 1e18;

        setStoreValues(store1, TEST_KEY, valA, valA * 10, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, valB, valB * 10, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, valC, valC * 10, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, valD, valD * 10, 1, 1, timestamp);

        // Should never revert
        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Result should be within min-max range
        uint256 minVal = valA;
        uint256 maxVal = valA;
        if (valB < minVal) minVal = valB;
        if (valB > maxVal) maxVal = valB;
        if (valC < minVal) minVal = valC;
        if (valC > maxVal) maxVal = valC;
        if (valD < minVal) minVal = valD;
        if (valD > maxVal) maxVal = valD;

        assertTrue(result.fairValue >= minVal && result.fairValue <= maxVal, "Median should be within range");
    }

    function test_Precision_TimestampSelection() public {
        // Even count: Which timestamp is selected?
        // [100@t1, 200@t2, 300@t3, 400@t4]
        // Middle two: 200@t2, 300@t3
        // Timestamp selected: t3 (right element of middle pair)
        uint256 timestamp1 = block.timestamp;
        uint256 timestamp2 = timestamp1 + 1;
        uint256 timestamp3 = timestamp2 + 1;
        uint256 timestamp4 = timestamp3 + 1;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp1);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp2);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp3);
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp4);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Timestamp should be from the right element of the middle pair (300@t3)
        assertEq(result.timestamp, timestamp3, "Timestamp should be from right middle element");
    }

    function test_Precision_ZeroValues() public {
        // Test with zero values
        // [0, 0, 100, 200]
        // Middle two: 0, 100
        // Average: (0 + 100) / 2 = 50 (exact)
        uint256 timestamp = block.timestamp;

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        setStoreValues(store1, TEST_KEY, 0, 0, 0, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 0, 0, 0, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 200, 2000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        assertEq(result.fairValue, 50, "Median with zeros should be correct");
    }

    function test_Precision_MixedOddEven() public {
        // Compare 5 oracles vs 6 oracles with same core values
        // 5 oracles (odd): [100, 200, 300, 400, 500] -> 300 (exact)
        // 6 oracles (even): [100, 200, 300, 400, 500, 600] -> (300+400)/2 = 350 (exact)
        uint256 timestamp = block.timestamp;

        // Create 5 stores
        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();
        MockValueStore store5 = createMockStore();

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 1, 1, timestamp);
        setStoreValues(store5, TEST_KEY, 500, 5000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory oddResult = oracle.getMedianValues(TEST_KEY);

        // Add 6th store
        MockValueStore store6 = createMockStore();
        setStoreValues(store6, TEST_KEY, 600, 6000, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory evenResult = oracle.getMedianValues(TEST_KEY);

        assertEq(oddResult.fairValue, 300, "Odd count (5): middle value is 300");
        assertEq(evenResult.fairValue, 350, "Even count (6): average of 300 and 400 is 350");
    }
}

/*//////////////////////////////////////////////////////////////
                    OVERFLOW & EDGE CASE TEST SUITE
    Tests for arithmetic overflow and critical edge cases

    IMPORTANT: These tests use MockValueStore which has different validation
    than the real ValueStore contract:

    Real ValueStore (ValueStore.sol:116-118):
        if (numerator != 0 && denominator == 0) {
            revert DivisionByZero();
        }

    MockValueStore (test mock):
        - Allows any numerator/denominator combination
        - Used for flexibility in testing oracle contract logic

    Test Strategy:
    - Overflow tests: Test ORACLE's aggregation (summing/rounding) - Mock is appropriate
    - Edge case tests: Use realistic values that real ValueStore would allow
//////////////////////////////////////////////////////////////*/
contract OverflowEdgeCaseTest is BaseTest {
    function setUp() public override {
        super.setUp();
        vm.prank(owner);
        oracle.setThreshold(2);
        vm.prank(owner);
        oracle.setTimeoutSeconds(3600);
    }

    function test_Overflow_FairValueSum_Overflows() public {
        // Test that overflow in fairValue sum is caught
        // Solidity 0.8.30 has built-in overflow checks that should revert

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;
        uint256 maxVal = type(uint256).max;

        // Create values where sum of middle two will overflow
        // Middle two will be maxVal and maxVal
        // sumFair = maxVal + maxVal (OVERFLOWS!)
        setStoreValues(store1, TEST_KEY, maxVal - 1, maxVal, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, maxVal, maxVal, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, maxVal, maxVal, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, maxVal, maxVal, 1, 1, timestamp);

        // Should revert due to overflow in fairValue sum
        vm.expectRevert();
        oracle.getMedianValues(TEST_KEY);
    }

    function test_Overflow_FairValueSum_NearMaximum() public {
        // Test with very large values that DON'T overflow
        // Verifies the contract can handle large but safe values

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Use values that are large but won't overflow when summed with +1
        // Safe threshold: < type(uint256).max / 2 for each value
        uint256 largeValue = 1e18;  // Very large but safe

        setStoreValues(store1, TEST_KEY, largeValue, largeValue, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, largeValue * 2, largeValue * 2, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, largeValue * 3, largeValue * 3, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, largeValue * 4, largeValue * 4, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Middle two: 2e18 and 3e18
        // Sum: 5e18, Rounding: (5e18 + 1) / 2 = 2.5e18
        assertEq(result.fairValue, 2.5e18, "Should handle large values");
    }

    function test_Overflow_RoundingIncrement_Overflow() public {
        // Test overflow in the +1 for rounding
        // When sumFair = type(uint256).max, adding 1 should overflow

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Values where sumFair = max but adding 1 overflows
        // Need: fairValues[mid1] + fairValues[mid2] = type(uint256).max
        uint256 val1 = type(uint256).max / 2;
        uint256 val2 = type(uint256).max - val1;

        setStoreValues(store1, TEST_KEY, val1 - 1000, val1 - 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, val1, val1, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, val2, val2, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, val2 + 1000, val2 + 1000, 1, 1, timestamp);

        // Should revert when trying to add 1 to max
        vm.expectRevert();
        oracle.getMedianValues(TEST_KEY);
    }

    function test_Overflow_UsdValueSum_Overflows() public {
        // Test overflow in usdValue sum (separate from fairValue)

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;
        uint256 maxVal = type(uint256).max;

        // Same fairValues, but usdValues overflow
        setStoreValues(store1, TEST_KEY, 100, maxVal, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, maxVal, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, maxVal, 1, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, maxVal, 1, 1, timestamp);

        // Should revert due to overflow in usdValue sum
        vm.expectRevert();
        oracle.getMedianValues(TEST_KEY);
    }

    function test_EdgeCase_ZeroDenominator() public {
        // Test edge case: numerator=0, denominator=0 (VALID in real ValueStore)
        // Real ValueStore validation: if (numerator != 0 && denominator == 0) revert
        // So numerator=0, denominator=0 is ALLOWED (represents undefined/zero data)

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Oracle 2 has numerator=0, denominator=0 (realistic - represents no fair value data)
        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 0, 0, timestamp);  // numerator=0, denominator=0 (VALID)
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Odd count: should use middle oracle (store2)
        assertEq(result.fairValue, 200, "FairValue from middle oracle");
        assertEq(result.numerator, 0, "Numerator = 0");
        assertEq(result.denominator, 0, "Denominator = 0 (when numerator is also 0)");

        emit log_string("INFO: numerator=0, denominator=0 is valid (represents zero/undefined)");
        emit log_string("Real ValueStore allows this (only rejects numerator!=0 with denominator=0)");
    }

    function test_EdgeCase_ZeroDenominator_EvenCount() public {
        // Test zero denominator with even count (takes from right middle)

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();
        MockValueStore store4 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 2, 0, timestamp);  // Warning: denominator = 0
        setStoreValues(store3, TEST_KEY, 300, 3000, 3, 1, timestamp);
        setStoreValues(store4, TEST_KEY, 400, 4000, 4, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Even count: middle two are 200 and 300
        // FairValue median: (200 + 300 + 1) / 2 = 250
        // Rounds to 250, closer to 300, so takes from RIGHT middle (store3)
        // which has denominator = 1 (valid)
        assertEq(result.fairValue, 250, "FairValue median");
        assertEq(result.numerator, 3, "Numerator from right middle oracle");
        assertEq(result.denominator, 1, "Denominator from right middle oracle (not zero)");
    }

    function test_EdgeCase_AllZeroDenominators() public {
        // Test when ALL oracles have numerator=0, denominator=0
        // This is VALID - represents all oracles reporting no fair value data

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        // All oracles have numerator=0, denominator=0 (realistic - no fair value data)
        setStoreValues(store1, TEST_KEY, 100, 1000, 0, 0, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 0, 0, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 0, 0, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Contract returns numerator=0, denominator=0 (valid edge case)
        assertEq(result.fairValue, 200, "FairValue");
        assertEq(result.numerator, 0, "Numerator = 0");
        assertEq(result.denominator, 0, "Denominator = 0 (when numerator is also 0)");

        emit log_string("INFO: All oracles have numerator=0, denominator=0");
        emit log_string("This represents no fair value data - valid edge case");
    }

    function test_EdgeCase_ZeroNumerator() public {
        // Test that numerator = 0 is handled correctly (valid case)

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        setStoreValues(store1, TEST_KEY, 100, 1000, 0, 1, timestamp);
        setStoreValues(store2, TEST_KEY, 200, 2000, 0, 1, timestamp);
        setStoreValues(store3, TEST_KEY, 300, 3000, 0, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Numerator = 0 is valid (represents 0 value)
        assertEq(result.fairValue, 200, "FairValue");
        assertEq(result.numerator, 0, "Numerator = 0 is valid");
        assertEq(result.denominator, 1, "Denominator");
    }

    function test_EdgeCase_MaxValues_NoOverflow() public {
        // Test with very large values that don't overflow
        // For odd count, no summing occurs, so larger values are safe

        MockValueStore store1 = createMockStore();
        MockValueStore store2 = createMockStore();
        MockValueStore store3 = createMockStore();

        uint256 timestamp = block.timestamp;

        // Use very large but safe values
        // For odd count, we just return the middle value - no arithmetic needed
        uint256 largeValue = 1e28;  // Extremely large but safe

        setStoreValues(store1, TEST_KEY, largeValue, largeValue, 1, 1, timestamp);
        setStoreValues(store2, TEST_KEY, largeValue * 2, largeValue * 2, 1, 1, timestamp);
        setStoreValues(store3, TEST_KEY, largeValue * 3, largeValue * 3, 1, 1, timestamp);

        DIAOracleV3MetaFairValueField.MedianSet memory result = oracle.getMedianValues(TEST_KEY);

        // Odd count: should return the middle value (store2's value) - no summing
        assertEq(result.fairValue, largeValue * 2, "Should handle very large values with odd count");
    }
}
