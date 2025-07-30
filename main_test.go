package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_generateKeys(t *testing.T) {
	type args struct {
		mnemonic   string
		passphrase string
	}
	tests := []struct {
		name    string
		args    args
		want    *KeyResult
		wantErr bool
	}{
		{
			name: "standard 12-word mnemonic with passphrase",
			args: args{
				mnemonic:   "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
				passphrase: "test123",
			},
			want: &KeyResult{
				MasterPrivateKey: "xprv9s21ZrQH143K2tCU97YegdvoaLQUTP9QaHuXeB7ZZVGei4sFq9eoisaxBckjZhz6R7bxWrMQBJNTm9UrBVqu1fs4QJGinVMZiXcJxqRfNMy",
				MasterPublicKey:  "xpub661MyMwAqRbcFNGwF95f3msY8NExrqsFwWq8SZXB7podasCQNgy4GfuS2uS5iPsoNNicMX7WfeDtHJybZPAJZvM36iNGUxc1JVEEST152xu",
				SSHPrivateKey: `-----BEGIN OPENSSH PRIVATE KEY-----
OeHaM1Lz7xeBiM/Pdy3V+oF3xRwzDZybRUhWG5dzyHv8O8x/1tyLkJYNshRqrKAc
5aTlg9hSt0xc/Nu7zvuBzA==
-----END OPENSSH PRIVATE KEY-----
`,
				SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPw7zH/W3IuQlg2yFGqsoBzlpOWD2FK3TFz827vO+4HM\n",
			},
			wantErr: false,
		},
		{
			name: "standard 12-word mnemonic without passphrase",
			args: args{
				mnemonic:   "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
				passphrase: "",
			},
			want: &KeyResult{
				MasterPrivateKey: "xprv9s21ZrQH143K3GJpoapnV8SFfukcVBSfeCficPSGfubmSFDxo1kuHnLisriDvSnRRuL2Qrg5ggqHKNVpxR86QEC8w35uxmGoggxtQTPvfUu",
				MasterPublicKey:  "xpub661MyMwAqRbcFkPHucMnrGNzDwb6teAX1RbKQmqtEF8kK3Z7LZ59qafCjB9eCRLiTVG3uxBxgKvRgbubRhqSKXnGGb1aoaqLrpMBDrVxga8",
				SSHPrivateKey: `-----BEGIN OPENSSH PRIVATE KEY-----
GDfBvo4plewRzaKwZhUb4s+0it+eR7FR1Gras6Ic32ci1H1KEiv/ccU9o+TyAmBy
y23q89I0DyQdmEX2bPUoVg==
-----END OPENSSH PRIVATE KEY-----
`,
				SSHPublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICLUfUoSK/9xxT2j5PICYHLLberz0jQPJB2YRfZs9ShW\n",
			},
			wantErr: false,
		},
		{
			name: "different 12-word mnemonic",
			args: args{
				mnemonic:   "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong",
				passphrase: "password",
			},
			want:    nil, // Will be verified by format checks only
			wantErr: false,
		},
		{
			name: "24-word mnemonic",
			args: args{
				mnemonic:   "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art",
				passphrase: "test",
			},
			want:    nil, // Will be verified by format checks only
			wantErr: false,
		},
		{
			name: "complex passphrase",
			args: args{
				mnemonic:   "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
				passphrase: "complex!@#$%^&*()_+passphrase",
			},
			want:    nil, // Will be verified by format checks only
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateKeys(tt.args.mnemonic, tt.args.passphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			// Verify basic structure
			if got.MasterPrivateKey == "" {
				t.Error("MasterPrivateKey is empty")
			}
			if got.MasterPublicKey == "" {
				t.Error("MasterPublicKey is empty")
			}
			if got.SSHPrivateKey == "" {
				t.Error("SSHPrivateKey is empty")
			}
			if got.SSHPublicKey == "" {
				t.Error("SSHPublicKey is empty")
			}

			// Verify SSH private key format
			if !strings.Contains(got.SSHPrivateKey, "-----BEGIN OPENSSH PRIVATE KEY-----") {
				t.Error("SSH private key doesn't have proper PEM header")
			}
			if !strings.Contains(got.SSHPrivateKey, "-----END OPENSSH PRIVATE KEY-----") {
				t.Error("SSH private key doesn't have proper PEM footer")
			}

			// Verify SSH public key format
			if !strings.HasPrefix(got.SSHPublicKey, "ssh-ed25519 ") {
				t.Error("SSH public key doesn't start with ssh-ed25519")
			}

			// For specific test cases with expected values, verify exact matches
			if tt.want != nil {
				if tt.want.MasterPrivateKey != "" && got.MasterPrivateKey != tt.want.MasterPrivateKey {
					t.Errorf("MasterPrivateKey = %v, want %v", got.MasterPrivateKey, tt.want.MasterPrivateKey)
				}
				if tt.want.MasterPublicKey != "" && got.MasterPublicKey != tt.want.MasterPublicKey {
					t.Errorf("MasterPublicKey = %v, want %v", got.MasterPublicKey, tt.want.MasterPublicKey)
				}
				if tt.want.SSHPrivateKey != "" && got.SSHPrivateKey != tt.want.SSHPrivateKey {
					t.Errorf("SSHPrivateKey = %v, want %v", got.SSHPrivateKey, tt.want.SSHPrivateKey)
				}
				if tt.want.SSHPublicKey != "" && got.SSHPublicKey != tt.want.SSHPublicKey {
					t.Errorf("SSHPublicKey = %v, want %v", got.SSHPublicKey, tt.want.SSHPublicKey)
				}
			}

			// Test deterministic generation - same inputs should produce same outputs
			got2, err2 := generateKeys(tt.args.mnemonic, tt.args.passphrase)
			if err2 != nil {
				t.Errorf("Second generateKeys() call failed: %v", err2)
				return
			}

			if !reflect.DeepEqual(got, got2) {
				t.Error("Key generation is not deterministic")
			}
		})
	}
}
