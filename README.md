# Mnemonic SSH

[![Go](https://github.com/pinpox/mnemonic-ssh/actions/workflows/go.yml/badge.svg)](https://github.com/pinpox/mnemonic-ssh/actions/workflows/go.yml)

A simple tool that generates deterministic SSH keypairs from mnemonic phrases
and passphrases using [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)/[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki) standards.

> [!NOTE]
> This tool was primarily made for my personal use. I'm sure there are other
> tools that can already do this, but I wanted something simple enough to be
> fully understood easily.

## Features

- Generate SSH ED25519 keypairs from [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) mnemonic phrases
- Support for optional passphrases for additional security
- Deterministic key generation - same mnemonic + passphrase always produces the same keys
- Compatible with standard [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) word lists
- Supports both 12-word and 24-word mnemonics
- Outputs SSH keys in standard OpenSSH format

## Usage

Run the program and follow the interactive prompts:

```bash
./mnemonic-ssh
```

The program will ask for:
1. **Passphrase** (optional) - Additional security layer
2. **Mnemonic words** - Either provide your own or leave empty to generate new ones

### Example Output

```
Enter passphrase (or leave empty for none):
test123
Enter mnemonic words (or leave empty to generate):
word1 word2 word3 word4 word5 word6 word7 word8 word9 word10 word11 word12
Mnemonic: 'word1 word2 word3 word4 word5 word6 word7 word8 word9 word10 word11 word12'
Passphrase: 'test123'
Master private key:  xprv9s21ZrQH143K2a8A5MyYweJ1TxU5YXBiSsoBTnwDQEydCmJSgJC2Kg6K46dEBUC5Bj3PP3f7k6h1ZNnaNHSwwThuiTF2jSH3CgmhndJ129Z
Master public key:  xpub661MyMwAqRbcF4CdBPWZJnEk1zJZwyuZp6inGBLpxaWc5ZdbDqWGsUQnuPPNPQhubpbvekCEJopkSU5hGAV6D5SPds88vtfbXZZqb5XaVtk

SSH Private Key:

-----BEGIN OPENSSH PRIVATE KEY-----
hiVGXbGTZDsMdg2YP0r1T7SsgteQ+qG0I7mQSY/JIDGkKnnYwfBGvF/tLACn+X3w
FewNI1B66Et3xcgHueEICQ==
-----END OPENSSH PRIVATE KEY-----

SSH Public Key:

ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKQqedjB8Ea8X+0sAKf5ffAV7A0jUHroS3fFyAe54QgJ
```

## Testing

Run the test suite:

```bash
go test -v
```

The tests verify:
- Key generation with various mnemonic/passphrase combinations
- Deterministic behavior (same inputs = same outputs)
- Proper SSH key formatting
- Expected values for known test vectors

## Use Cases

- **Backup Recovery**: Restore SSH keys from mnemonic backup
- **Multi-Device Sync**: Generate same SSH keys across different devices
- **Secure Key Storage**: Store SSH keys as human-readable mnemonic phrases
- **Deterministic Infrastructure**: Reproducible SSH key generation for automation
