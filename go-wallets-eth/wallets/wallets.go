package wallets

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"eth-wallet/models"
)

const (
	mnemonic = "mandate author pulp rhythm physical naive trigger supply ride truly raccoon parrot"
)

var count int = 0

func GetAddress() (models.Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	addr_index := strconv.Itoa(count)
	hdPath := "m/44'/60'/0'/0/" + addr_index
	fmt.Println(hdPath)
	addr, pub, pri, err := getKeys(wallet, hdPath)
	if err != nil {
		return models.Wallet{}, err
	}
	increment()
	return models.Wallet{
		PrivateKey: pri,
		PublicKey:  pub,
		Address:    addr,
		HDpath:     hdPath,
	}, nil
}
func increment() {
	count = count + 1
}
func getKeys(wallet *hdwallet.Wallet, hdpath string) (string, string, string, error) {
	path := hdwallet.MustParseDerivationPath(hdpath)
	accounts, err := wallet.Derive(path, false)
	if err != nil {
		return "", "", "", err
	}
	pvk, err := wallet.PrivateKey(accounts)
	if err != nil {
		return "", "", "", err
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
