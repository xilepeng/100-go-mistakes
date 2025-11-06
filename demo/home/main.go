package main

import "fmt"

func main() {
	var (
		interestRate float32 = 0.045    // 年化利率
		irincipal    float32 = 300_0000 // 本金 300万
		interest     float32            // 利息
	)

	interest = irincipal * interestRate
	fmt.Printf("定存=%f,年化利率=%f,一年利息=%f\n", irincipal, interestRate, interest)
	fmt.Println("利息可供月租=", interest/12)
}

// 定存=300万, 年化利率=4.5%, 一年利息=13.5万
// 利息可供月租= 1.125万


/*
本金 principal 利息 interest 利率 interest rate 单利simple interest

复利compound interest

法定利息legal interest 手续费service charge

活期存款 current ( demand) deposit

定期存款 fixed ( time) deposit

存折 deposit book/passbook/bankbook

存单 certificate of deposit

存款单 pay-in slip/deposit slip( form, ticket)

存款收据deposit receipt

户名 account name

账号 account number

开户日期 opening date

存款期deposit term

密码 password/code/pin number
*/
