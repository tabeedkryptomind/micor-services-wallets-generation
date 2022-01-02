package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const (
	mnemonic = "mandate author pulp rhythm physical naive trigger supply ride truly raccoon parrot"
)
type Wallet struct{
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
	HDpath     string `json:"hdpath"`
}
var count  int = 0
func main() {
	
	route := gin.Default()
	route.GET("/api/eth/create-wallet/", generateWallets)
	route.Run(":8080")
}
func getAddress()(Wallet, error){
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	addr_index := strconv.Itoa(count)
	hdPath := "m/44'/60'/0'/0/" + addr_index
	fmt.Println(hdPath)
	pri,pub,addr , err := getKeys(wallet, hdPath)
	if err != nil{
		return Wallet{}, err
	}
	increment()
	return Wallet{
		PrivateKey: pri,
		PublicKey: pub,
		Address: addr,
		HDpath: hdPath,
	}, nil
}

func getKeys(wallet *hdwallet.Wallet, hdpath string)(string, string, string, error){
	path := hdwallet.MustParseDerivationPath(hdpath)
	accounts, err := wallet.Derive(path, false)
	if err != nil {
		return  "", "", "", err
	}
	pvk, err := wallet.PrivateKey(accounts)
	if err != nil {
		return  "", "", "", err
	}
	// generate private key	
	pvbytes := crypto.FromECDSA(pvk)
	privKey := hexutil.Encode(pvbytes)
	// generate public key

	pubBytes := crypto.FromECDSAPub(&pvk.PublicKey)
	pubKey := hexutil.Encode(pubBytes)
	// address hex will give public address
	return accounts.Address.Hex(), pubKey, privKey, nil
}

func increment(){
	count = count +1 
}
func generateWallets(ctx * gin.Context){
	w , err :=  getAddress()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create wallet",
			"status": http.StatusInternalServerError,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"wallet": w,
	})
}