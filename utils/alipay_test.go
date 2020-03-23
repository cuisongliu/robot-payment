package utils

import (
	"testing"

	"github.com/smartwalle/alipay/v3"
)

func TestAlipay_PayTo(t *testing.T) {
	type fields struct {
		AppPrivateKey       string
		AppCertPublicKey    string
		AlipayRootCert      string
		AlipayCertPublicKey string
		AppId               string
		Client              *alipay.Client
	}
	type args struct {
		account string
		amount  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "pay fanux",
			args:    args{"15805691422", "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New()
			if err := a.PayTo(tt.args.account, tt.args.amount,"感谢完成xxx任务"); (err != nil) != tt.wantErr {
				t.Errorf("PayTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
