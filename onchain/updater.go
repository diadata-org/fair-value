package onchain

import (
	"context"
	"math"
	"math/big"
	"time"

	ValueStore "github.com/diadata-org/fair-value/contracts/valuestore"
	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

const (
	DECIMALS_ORACLE_VALUE = 8
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	loglevel, err := logrus.ParseLevel(utils.Getenv("LOG_LEVEL_UPDATER", "info"))
	if err != nil {
		log.Errorf("Parse log level: %v.", err)
	}
	log.SetLevel(loglevel)
}

func OracleUpdateExecutor(
	auth *bind.TransactOpts,
	contract *ValueStore.ValueStore,
	contractBackup *ValueStore.ValueStore,
	conn *ethclient.Client,
	connBackup *ethclient.Client,
	data []models.FairValueData,
) {

	for _, d := range data {
		timestamp := time.Now().Unix()
		var keys []string
		var fairValues, valueUsds, numerators, denominators []*big.Int
		var timestamps []int64

		log.Infof(
			"updater - data received at %v: %v -- %v -- %v.",
			time.Unix(timestamp, 0),
			d.Symbol,
			d.PriceUSD,
			d.Time,
		)

		key := d.Symbol
		keys = append(keys, key)
		fairValues = append(fairValues, big.NewInt(int64(d.FairValueNative*math.Pow10(int(DECIMALS_ORACLE_VALUE)))))
		valueUsds = append(valueUsds, big.NewInt(int64(d.PriceUSD*math.Pow10(int(DECIMALS_ORACLE_VALUE)))))
		numerators = append(numerators, d.Numerator)
		denominators = append(denominators, d.Denominator)

		timestamps = append(timestamps, d.Time.Unix())

		err := updateOracleMultiValues(conn, contract, auth, keys, fairValues, valueUsds, numerators, denominators, timestamps)
		if err != nil {
			log.Warnf("updater - Failed to update Oracle: %v. Retry with backup node.", err)
			err := updateOracleMultiValues(connBackup, contractBackup, auth, keys, fairValues, valueUsds, numerators, denominators, timestamps)
			if err != nil {
				log.Errorf("backup updater - Failed to update Oracle: %v.", err)
				return
			}
		}
	}
}

// TO DO: Either remove timestamps or rewrite contract such that every entry has its own timestamp.
func updateOracleMultiValues(
	client *ethclient.Client,
	contract *ValueStore.ValueStore,
	auth *bind.TransactOpts,
	keys []string,
	fairValues []*big.Int,
	valueUsds []*big.Int,
	numerators []*big.Int,
	denominators []*big.Int,
	timestamps []int64) error {

	// Get gas price suggestion
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("updater - SuggestGasPrice: %v.", err)
		return err
	}
	// Get 110% of the gas price
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(
		&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		},
		keys,
		fairValues,
		valueUsds,
		numerators,
		denominators,
	)
	if err != nil {
		return err
	}

	// log.Infof("updater - Gas price: %d.", tx.GasPrice())
	// log.Printf("Data: %x\n", tx.Data())
	// log.Infof("updater - Nonce: %d.", tx.Nonce())
	// log.Infof("updater - Tx To: %s.", tx.To().String())
	log.Infof("updater - Tx Hash: 0x%x.", tx.Hash())
	return nil
}
