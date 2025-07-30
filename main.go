package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {

	var mnemonic, passphrase string
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	// Ask for passphrase
	fmt.Println("Enter passphrase (or leave empty for none):")
	scanner.Scan()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	passphrase = scanner.Text()

	// Ask for mnemonic
	fmt.Println("Enter mnemonic words (or leave empty to generate):")
	scanner.Scan()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	mnemonic = scanner.Text()

	if len(mnemonic) == 0 {
		fmt.Println("No mnemonic provided")
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			panic(err)
		}

		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			panic(err)
		}
	}

	// Generate seed and keys
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Printf("Mnemonic: '%s'\n", mnemonic)
	fmt.Printf("Passphrase: '%s'\n", passphrase)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}
