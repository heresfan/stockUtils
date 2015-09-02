package main

import (
	"flag"
	"fmt"
	"os"
	"stockUtils/apis"
	"time"
)

var stockId *string = flag.String("sid", "sz002066,sz000048", "stock id")

func main() {
	flag.Parse()

	sid := *stockId
	var ret apis.BaiduStoctInfo

	for {
		if _, err := apis.BaiduStoctRequest(sid, &ret); err != nil {
			fmt.Printf("Failed err[%v]\n", err)
			os.Exit(1)
		}

		fmt.Printf("\033[2J")
		for _, v := range ret.Retdata.Stockinfo {
			fmt.Println("-------------------------------------------------------")
			fmt.Printf("timestamp		:%v\n", v.Time)
			fmt.Printf("name			:%v\n", v.Name)
			fmt.Printf("high price		:%v\n", v.Hprice)
			fmt.Printf("low price		:%v\n", v.Lprice)
			fmt.Printf("current price		:%v\n", v.Currentprice)
			fmt.Printf("growth percentage	:%.2f%% \n", (v.Currentprice/
				v.Closingprice-1)*100)

			fmt.Println("")
			fmt.Printf("sell 5	:	%v 		%v \n", v.Sellfiveprice, v.Sellfive/100)
			fmt.Printf("sell 4	:	%v 		%v \n", v.Sellfourprice, v.Sellfour/100)
			fmt.Printf("sell 3	:	%v 		%v \n", v.Sellthreeprice, v.Sellthree/100)
			fmt.Printf("sell 2	:	%v 		%v \n", v.Selltwoprice, v.Selltwo/100)
			fmt.Printf("sell 1	:	%v 		%v \n", v.Selloneprice, v.Sellone/100)
			fmt.Println("")
			fmt.Printf("buy  1	:	%v 		%v\n", v.Buyoneprice, v.Buyone/100)
			fmt.Printf("buy  2	:	%v 		%v\n", v.Buytwoprice, v.Buytwo/100)
			fmt.Printf("buy  3	:	%v 		%v\n", v.Buythreeprice, v.Buythree/100)
			fmt.Printf("buy  4	:	%v 		%v\n", v.Buyfourprice, v.Buyfour/100)
			fmt.Printf("buy  5	:	%v 		%v\n", v.Buyfiveprice, v.Buyfive/100)
			fmt.Printf("\n")
		}

		fmt.Printf("------------------------- %v -----------------------------\n", ret.Retdata.Market.Shanghai.Name)
		fmt.Printf("current dot		: %v\n", ret.Retdata.Market.Shanghai.Curdot)
		fmt.Printf("current price		: %v\n", ret.Retdata.Market.Shanghai.Curprice)
		fmt.Printf("growth rate		: %v%%\n", ret.Retdata.Market.Shanghai.Rate)
		//fmt.Printf("deal number     : %v\n", ret.Retdata.Market.Shanghai.Dealnumber)
		fmt.Printf("turnover		: %v\n", ret.Retdata.Market.Shanghai.Turnover)

		time.Sleep(1 * time.Second)
	}

	os.Exit(0)
}
