package utils

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"os"
	"time"
)

type Alipay struct {
	AppPrivateKey       string
	AppCertPublicKey    string
	AlipayRootCert      string
	AlipayCertPublicKey string
	AppId string
	Client *alipay.Client
}

func New() *Alipay {
	a := &Alipay{
		AppPrivateKey:      os.Getenv("AppPrivateKey"),
		AppCertPublicKey:    os.Getenv("AppCertPublicKey"),
		AlipayRootCert:      os.Getenv("AlipayRootCert"),
		AlipayCertPublicKey: os.Getenv("AlipayCertPublicKey"),
		AppId:               os.Getenv("AppId"),
	}

	fmt.Println()
	fmt.Println(a.AppPrivateKey)
	fmt.Println(a.AppCertPublicKey)
	fmt.Println(a.AlipayRootCert)
	fmt.Println(a.AlipayCertPublicKey)
	fmt.Println(a.AppId)
	fmt.Println()

	client,err := alipay.New(a.AppId,a.AppPrivateKey,true)
	if err != nil {
		fmt.Printf("create client failed %s",err)
		return nil
	}
	client.LoadAppPublicCert(a.AppCertPublicKey)
	client.LoadAliPayRootCert(a.AlipayRootCert)
	client.LoadAliPayPublicCert(a.AlipayCertPublicKey)
	a.Client = client
	return a
}

// pay amount yuan to alipay account.
func (a *Alipay) PayTo(account string, amount string, remark string) error {
	p := alipay.FundTransToAccountTransfer{
		AppAuthToken:  "",
		OutBizNo:  fmt.Sprintf("%d-%s-%s",time.Now().Unix(),account,amount),
		PayeeType:     "ALIPAY_LOGONID",
		PayeeAccount:  account,
		Amount:       amount,
		PayerShowName: "sealyun",
		PayeeRealName: "",
		Remark:        remark,
	}
	fmt.Printf("out_biz_no : %s",p.OutBizNo)
	rsp, err := a.Client.FundTransToAccountTransfer(p)
	if err != nil {
		return fmt.Errorf("pay for %s failed %s, resp : %v",account,err,rsp)
	}
	if !rsp.IsSuccess(){
		return fmt.Errorf("pay for %s not success %s, %v",account,err,rsp)
	}
	return nil
}
