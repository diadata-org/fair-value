package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	ANKR_BASE_URL    = "https://rpc.ankr.com/premium-http/btc_blockbook/"
	MEMPOOL_BASE_URL = "https://mempool.space/api/"
)

type ankrAPIResponse struct {
	Balance            string `json:"balance"`
	TotalReceived      string `json:"totalReceived"`
	TotalSent          string `json:"totalSent"`
	UnconfirmedBalance string `json:"unconfirmedBalance"`
}

type mempoolAPIResponse struct {
	ChainStats mempoolChainStats `json:"chain_stats"`
}

type mempoolChainStats struct {
	FundedTxoSum int `json:"funded_txo_sum"`
	SpentTxoSum  int `json:"spent_txo_sum"`
}

// RPC request payload
type RPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

// RPC response for scantxoutset
type ScantxoutsetResponse struct {
	Result struct {
		TotalAmount float64 `json:"total_amount"`
	} `json:"result"`
	Error any `json:"error"`
}

func GetBitcoinWalletBalance(walletAddress string, bitcoinRPC string) (float64, error) {

	// Prepare scantxoutset RPC payload
	payload := RPCRequest{
		Jsonrpc: "1.0",
		ID:      "curltest",
		Method:  "scantxoutset",
		Params:  []any{"start", []string{fmt.Sprintf("addr(%s)", walletAddress)}},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", bitcoinRPC, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("pBTC -- Bitcoin RPC http error, i.e. status code %v for wallet %s", resp.StatusCode, walletAddress)
		return 0, err
	}

	var rpcResp ScantxoutsetResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return 0, err
	}

	if rpcResp.Error != nil {
		return 0, err
	}

	return rpcResp.Result.TotalAmount, nil
}

func GetBitcoinWalletBalanceAnkr(wallet string) (balance float64, err error) {
	url := ANKR_BASE_URL +
		Getenv("ANKR_BITCOIN_API_KEY", "") +
		"/api/v2/address/" + wallet

	data, _, err := GetRequest(url)
	if err != nil {
		return
	}

	var ar ankrAPIResponse
	err = json.Unmarshal(data, &ar)
	if err != nil {
		return
	}

	balanceInt, err := strconv.Atoi(ar.Balance)
	if err != nil {
		return
	}
	balance = float64(balanceInt) / 1e8
	return

}

func GetBitcoinWalletBalanceMempool(wallet string) (balance float64, err error) {
	url := MEMPOOL_BASE_URL + "address/" + wallet

	data, _, err := GetRequest(url)
	if err != nil {
		return
	}

	var mr mempoolAPIResponse
	err = json.Unmarshal(data, &mr)
	if err != nil {
		return
	}

	balance += float64(mr.ChainStats.FundedTxoSum-mr.ChainStats.SpentTxoSum) / 1e8

	return

}
