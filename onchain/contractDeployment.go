package onchain

import (
	"encoding/hex"
	"os"
	"strings"
	"time"

	erc1967proxy "github.com/diadata-org/fair-value/contracts/erc1967proxy"
	ValueStore "github.com/diadata-org/fair-value/contracts/valuestore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployOrBindContract(
	deployedContract string,
	conn *ethclient.Client,
	connBackup *ethclient.Client,
	auth *bind.TransactOpts,
	contract **ValueStore.ValueStore,
	contractBackup **ValueStore.ValueStore,
) error {
	var err error
	if deployedContract != "" {

		// bind primary and backup
		*contract, err = ValueStore.NewValueStore(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
		*contractBackup, err = ValueStore.NewValueStore(common.HexToAddress(deployedContract), connBackup)
		if err != nil {
			return err
		}

	} else {
		var implAddr common.Address
		var implTx *types.Transaction
		implAddr, implTx, *contract, err = ValueStore.DeployValueStore(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy ValueStore implementation: %v", err)
			return err
		}
		log.Infof("ValueStore Implementation pending deploy: 0x%x.", implAddr)
		log.Infof("Implementation Transaction waiting to be mined: 0x%x.", implTx.Hash())

		log.Info("Waiting for implementation deployment to be mined...")
		time.Sleep(30 * time.Second)

		proxyBytecodeHex, err := os.ReadFile("/tmp/ERC1967Proxy.bin")
		if err != nil {
			log.Fatalf("could not read proxy bytecode: %v", err)
			return err
		}

		hexString := strings.TrimSpace(strings.TrimPrefix(string(proxyBytecodeHex), "0x"))
		proxyBytecode, err := hex.DecodeString(hexString)
		if err != nil {
			log.Fatalf("could not decode proxy bytecode: %v", err)
			return err
		}

		proxyABI, err := erc1967proxy.ERC1967ProxyMetaData.GetAbi()
		if err != nil {
			log.Fatalf("could not parse proxy ABI: %v", err)
			return err
		}

		valueStoreABI, err := ValueStore.ValueStoreMetaData.GetAbi()
		if err != nil {
			log.Fatalf("could not parse ValueStore ABI: %v", err)
			return err
		}

		initializeData, err := valueStoreABI.Pack("initialize", auth.From)
		if err != nil {
			log.Fatalf("could not pack initialize data: %v", err)
			return err
		}

		proxyAddr, proxyTx, _, err := bind.DeployContract(
			auth,
			*proxyABI,
			proxyBytecode,
			conn,
			implAddr,
			initializeData,
		)
		if err != nil {
			log.Fatalf("could not deploy proxy: %v", err)
			return err
		}

		log.Infof("ERC1967Proxy pending deploy: 0x%x.", proxyAddr)
		log.Infof("Proxy Transaction waiting to be mined: 0x%x.", proxyTx.Hash())

		*contract, err = ValueStore.NewValueStore(proxyAddr, conn)
		if err != nil {
			log.Fatalf("could not bind to proxy: %v", err)
			return err
		}

		*contractBackup, err = ValueStore.NewValueStore(proxyAddr, connBackup)
		if err != nil {
			log.Fatalf("could not bind backup to proxy: %v", err)
			return err
		}

		log.Info("Waiting for proxy deployment to be mined...")
		time.Sleep(30 * time.Second)

		owner, err := (*contract).Owner(nil)
		if err != nil {
			log.Warnf("Could not verify owner: %v", err)
		} else {
			log.Infof("Proxy owner verified: 0x%x", owner)
			if owner != auth.From {
				log.Warnf("Owner mismatch! Expected: 0x%x, Got: 0x%x", auth.From, owner)
			} else {
				log.Info("Owner correctly set to deployer address")
			}
		}

		log.Info("Deployment successful!")
		log.Infof("ValueStore Implementation: 0x%x", implAddr)
		log.Infof("ERC1967Proxy (use this address): 0x%x", proxyAddr)
		log.Infof("Deployer (Owner): 0x%x", auth.From)
	}
	return nil
}
