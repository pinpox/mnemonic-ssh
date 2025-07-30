package main

import (
	"bufio"
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh"
)

type KeyResult struct {
	MasterPrivateKey string
	MasterPublicKey  string
	SSHPrivateKey    string
	SSHPublicKey     string
}

func generateKeys(mnemonic, passphrase string) (*KeyResult, error) {
	// Generate seed and keys
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	publicKey := masterKey.PublicKey()

	// Generate SSH keypair from master key
	// Use the master key bytes as seed for ED25519 key generation
	sshPrivateKey := ed25519.NewKeyFromSeed(masterKey.Key[:32])
	sshPublicKey := sshPrivateKey.Public().(ed25519.PublicKey)

	// Convert to SSH format
	sshPubKey, err := ssh.NewPublicKey(sshPublicKey)
	if err != nil {
		return nil, err
	}

	// Format SSH private key (using OpenSSH format)
	privKeyPEM := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: sshPrivateKey,
	}

	return &KeyResult{
		MasterPrivateKey: masterKey.String(),
		MasterPublicKey:  publicKey.String(),
		SSHPrivateKey:    string(pem.EncodeToMemory(privKeyPEM)),
		SSHPublicKey:     string(ssh.MarshalAuthorizedKey(sshPubKey)),
	}, nil
}

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

	// Generate keys
	result, err := generateKeys(mnemonic, passphrase)
	if err != nil {
		panic(err)
	}

	// Display mnemonic and keys
	fmt.Printf("Mnemonic: '%s'\n", mnemonic)
	fmt.Printf("Passphrase: '%s'\n", passphrase)
	fmt.Println("Master private key: ", result.MasterPrivateKey)
	fmt.Println("Master public key: ", result.MasterPublicKey)
	fmt.Println()
	fmt.Println("SSH Private Key:")
	fmt.Println()
	fmt.Printf("%s", result.SSHPrivateKey)

	fmt.Println()
	fmt.Println("SSH Public Key:")
	fmt.Println()

	fmt.Printf("%s", result.SSHPublicKey)
}
