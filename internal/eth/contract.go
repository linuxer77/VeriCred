package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"vericred/internal/eth/build"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


func Contractresp(option string) {

	privateKeyHex := "e27fc11f6468ca7b5916ebfd853b0e92aca9a5899af26414bae8d22bfd55c5d4"
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatal("Invalid private key")
	}

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/2a159ca7304a4df4ade0d0ccf9ce70ef")
	if err != nil {
		log.Fatal("❌ Failed to connect to the sepolia network.", err)
	}

	log.Println("Successfully logged in to sepolia network")

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	log.Println("contex.Background():", context.Background())
	chainID, _ := client.NetworkID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.From = fromAddress
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)


	address, tx, instance, err := build.DeployBuild(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy:", err)
	}

	fmt.Println("Contract deployed at: ", address.Hex())
	fmt.Println("Deployment TX:", tx.Hash().Hex())

	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Wait failed: ", err)
	}
	
	// cAddress := common.HexToAddress("0x2a67320e4fc9d8E072021575e5D0450ED5ae316e")

	// instance, err = build.NewBuild(cAddress, client)
	// if err != nil {
	// 	log.Fatal("❌ Failed to bind to contract:", err)
	// }
	// log.Println("Instance: ", instance)
	
	userAddr := common.HexToAddress("0x2a67320e4fc9d8E072021575e5D0450ED5ae316e")
	_, err = instance.MintDoc(auth, userAddr, "https://chocolate-electoral-parrot-267.mypinata.cloud/ipfs/bafkreifgophqzx5pw2uvukbozh7fxcpo5fdx2rprh73t3l72pgszekeita")
	if err != nil {
		log.Fatal("❌ Failed to call AllOrgs:", err)
	}

	tokenId, err := instance.CurrentTokenId(nil)
	if err != nil {
		log.Fatal("Failed to call currenttokenid: ", err)
	}

	fmt.Println("TokenID:", tokenId)
}

