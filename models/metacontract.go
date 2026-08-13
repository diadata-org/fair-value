package models

import (
	"strconv"

	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

func MakeMetacontract(address string, precision string, rpcURL, rpcURLBackup string) (metacontract MetacontractData, err error) {
	metacontract.Address = common.HexToAddress(address)
	metacontract.Precision, err = strconv.Atoi(precision)
	if err != nil {
		log.Errorf("parse METACONTRACT_FV_PRECISION: %v. Fallback to 18", err)
		metacontract.Precision = 18
	}
	metacontract.Client, err = utils.MakeEthClient(rpcURL, rpcURLBackup)
	return
}
