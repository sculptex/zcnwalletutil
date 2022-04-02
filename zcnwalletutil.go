package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"github.com/0chain/gosdk/core/zcncrypto"
)

func main() {	
	var mnemonic string
	var clientid string
	var privkey string
	var pubkey string
    var doJSON bool
    flag.StringVar(&mnemonic, "mnemonic", "", "12/24 word seed phrase, encase in quotes (\"\") on command line")
    flag.BoolVar(&doJSON, "json", false, "output as JSON")
    flag.Parse()
	sigScheme := zcncrypto.NewSignatureScheme("bls0chain")
	if(len(mnemonic)>0) {
		if !zcncrypto.IsMnemonicValid(mnemonic) {
			fmt.Println("Error: Invalid mnemonic: ", mnemonic)
			return
		}
		wallet , _ := sigScheme.RecoverKeys(mnemonic)
		mnemonic = sigScheme.GetMnemonic()
		clientid = wallet.ClientID
		privkey = sigScheme.GetPrivateKey()
		pubkey = sigScheme.GetPublicKey()
	} else {
		wallet , _ := sigScheme.GenerateKeys()	
		mnemonic = sigScheme.GetMnemonic()
		clientid = wallet.ClientID
		privkey = sigScheme.GetPrivateKey()
		pubkey = sigScheme.GetPublicKey()
	}
	if(doJSON) {
		j := make(map[string]string)
		j["client_id"] = clientid
		j["mnemonic"] = mnemonic
		j["public_key"] = pubkey
		j["private_key"] = privkey
		b , _ := json.Marshal(j)
		fmt.Println(string(b))
	} else {
		fmt.Println("client_id: ", clientid)
		fmt.Println("mnemonic: ", mnemonic)
		fmt.Println("public_key: ", pubkey)
		fmt.Println("private_key: ", privkey)
	}
}
