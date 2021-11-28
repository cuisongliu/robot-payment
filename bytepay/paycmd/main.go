package main

import (
	"fmt"
	"github.com/fanux/robot/utils"
	"strconv"
	"time"
)

var UserAliaccountMap = map[string]string{"fanux":"15805691422","PatHoo":"13926139093",
	"cuisongliu":"912387319@qq.com","zhangguanzhang":"zhangguanzhang@qq.com",
	"uglyliu":"footprints19940807@163.com", "jinnzy":"839444083@qq.com", "svolence":"15726646803"}

func floattostr(input_num float64) string {

	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}


func main(){
	var amout float64
	bar := map[string]float64{
		"cuisongliu":0.25,
		"zhangguanzhang":0.15,
		"PatHoo":0.13,
	}
	_ = bar
	foo := map[string]float64{
		"fanux":0.001,
	}
	amout = 1466.63
 	pay := utils.New()
 	if pay == nil {
 		return
	}
 	for k,v := range foo {
 		account,ok := UserAliaccountMap[k]
 		if !ok {
 			fmt.Println("找不到账户")
 			continue
		}
		money := floattostr(amout * v)
		fmt.Printf("pay to [%s] [%s]", account,money)
 		err := pay.PayTo(account,money,"sealyun 月度奖励，感谢支持~")
 		if err != nil {
 			fmt.Println(err)
		}
	}
	time.Sleep(3 * time.Second)
}
