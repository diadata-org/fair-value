// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Minimal Ownable pattern
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }
}

abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    constructor(address initialOwner) {
        _owner = initialOwner;
        emit OwnershipTransferred(address(0), initialOwner);
    }

    modifier onlyOwner() {
        require(owner() == _msgSender(), "Not owner");
        _;
    }

    function owner() public view returns (address) {
        return _owner;
    }
}

// Interfaces
interface IERC4626 {
    function convertToAssets(uint256 shares) external view returns (uint256);
}

interface IUSDpPriceOracle {
    function getValue(string calldata key) external view returns (uint128, uint128);
}

contract SusdpVaultWithExternalPrice is Ownable {
    address public susdpVaultAddress;
    address public externalUsdPriceFeed;
    uint256 public constant QUERY_DECIMALS = 18;

    // Query keys
    string public USDP_QUERY_STRING;
    string public USDP_MARKET_QUERY_STRING;

    mapping(string => uint256) public values;

    constructor(
        address _initialOwner,
        address _susdpVaultAddress,
        address _externalUsdPriceFeed
    ) Ownable(_initialOwner) {
        susdpVaultAddress = _susdpVaultAddress;
        externalUsdPriceFeed = _externalUsdPriceFeed;
        USDP_QUERY_STRING = "USDp/USD";
        USDP_MARKET_QUERY_STRING = "USDp-market/USD";
    }

    function setUsdpQueryString(string memory newUsdpQueryString) external onlyOwner {
        USDP_QUERY_STRING = newUsdpQueryString;
    }

    function setUsdpMarketQueryString(string memory newUsdpMarketQueryString) external onlyOwner {
        USDP_MARKET_QUERY_STRING = newUsdpMarketQueryString;
    }

    /** 
     * Main getter: covers all relevant keys.
     */
    function getValue(string memory key) public view returns (uint128, uint128) {
        if (_compareStrings(key, "fairValueUSD:USDp")) {
            return IUSDpPriceOracle(externalUsdPriceFeed).getValue("fairValueUSD:USDp");
        } else if (_compareStrings(key, USDP_QUERY_STRING)) {
            return (uint128(_getSusdpPrice()), uint128(block.timestamp));
        } else if (_compareStrings(key, USDP_MARKET_QUERY_STRING)) {
            return (uint128(_getSusdpMarketBasedPrice()), uint128(block.timestamp));
        } else {
            // TO DO: Can be removed? Or do we need to write USDp market price? Then we need a setter as well.
            uint256 cValue = values[key];
            uint128 timestamp = uint128(cValue % 2**128);
            uint128 value = uint128(cValue >> 128);
            return (value, timestamp);
        }
    }

    /** sUSDp price (internal helper) */
    function _getSusdpPrice() internal view returns (uint256) {
        IERC4626 susdpVaultInstance = IERC4626(susdpVaultAddress);
        (uint128 usdpPrice, ) = IUSDpPriceOracle(externalUsdPriceFeed).getValue("fairValueUSD:USDp");
        uint256 susdpRate = susdpVaultInstance.convertToAssets(10 ** QUERY_DECIMALS);
        return (susdpRate * usdpPrice) / (10 ** QUERY_DECIMALS);
    }

    /** sUSDp market-based price (internal helper) */
    function _getSusdpMarketBasedPrice() internal view returns (uint256) {
        IERC4626 susdpVaultInstance = IERC4626(susdpVaultAddress);
        // TO DO: Where to get USDp market price?
        (uint128 usdpMarketPrice, ) = IUSDpPriceOracle(externalUsdPriceFeed).getValue("USDp");
        uint256 susdpRate = susdpVaultInstance.convertToAssets(10 ** QUERY_DECIMALS);
        return (susdpRate * usdpMarketPrice) / (10 ** QUERY_DECIMALS);
    }

    /** String equality helper */
    function _compareStrings(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}