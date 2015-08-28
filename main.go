package main

import (
	"flag"
	"fmt"
	"os"
	"stockUtils/apis"
	"time"
)

var stockId *string = flag.String("sid", "sz002066", "stock id")

func main() {
	sid := *stockId
	var ret apis.BaiduStoctInfo

	for {
		if _, err := apis.BaiduStoctRequest(sid, &ret); err != nil {
			fmt.Printf("Failed err[%v]\n", err)
			os.Exit(1)
		}

		fmt.Printf("\033[2J")
		fmt.Printf("time:%v\n", ret.Retdata.Stockinfo[0].Time)
		fmt.Printf("name:%v\n", ret.Retdata.Stockinfo[0].Name)
		fmt.Printf("high price:%v\n", ret.Retdata.Stockinfo[0].Hprice)
		fmt.Printf("low price:%v\n", ret.Retdata.Stockinfo[0].Lprice)
		fmt.Printf("current price:%v\n", ret.Retdata.Stockinfo[0].Currentprice)
		fmt.Printf("growth percentage :%.2f%% \n", (ret.Retdata.Stockinfo[0].Currentprice/
			ret.Retdata.Stockinfo[0].Closingprice-1)*100)

		fmt.Println("-------------------------------------------------------")
		fmt.Printf("sell 5	:	%v 		%v \n", ret.Retdata.Stockinfo[0].Selloneprice, ret.Retdata.Stockinfo[0].Sellone)
		fmt.Printf("sell 4	:	%v 		%v \n", ret.Retdata.Stockinfo[0].Selltwoprice, ret.Retdata.Stockinfo[0].Selltwo)
		fmt.Printf("sell 3	:	%v 		%v \n", ret.Retdata.Stockinfo[0].Sellthreeprice, ret.Retdata.Stockinfo[0].Sellthree)
		fmt.Printf("sell 2	:	%v 		%v \n", ret.Retdata.Stockinfo[0].Sellfourprice, ret.Retdata.Stockinfo[0].Sellfour)
		fmt.Printf("sell 1	:	%v 		%v \n", ret.Retdata.Stockinfo[0].Sellfiveprice, ret.Retdata.Stockinfo[0].Sellfive)
		fmt.Printf("buy 1	:	%v 		%v\n", ret.Retdata.Stockinfo[0].Buyoneprice, ret.Retdata.Stockinfo[0].Buyone)
		fmt.Printf("buy 2	:	%v 		%v\n", ret.Retdata.Stockinfo[0].Buytwoprice, ret.Retdata.Stockinfo[0].Buytwo)
		fmt.Printf("buy 3	:	%v 		%v\n", ret.Retdata.Stockinfo[0].Buythreeprice, ret.Retdata.Stockinfo[0].Buythree)
		fmt.Printf("buy 4	:	%v 		%v\n", ret.Retdata.Stockinfo[0].Buyfourprice, ret.Retdata.Stockinfo[0].Buyfour)
		fmt.Printf("buy 5	:	%v 		%v\n", ret.Retdata.Stockinfo[0].Buyfiveprice, ret.Retdata.Stockinfo[0].Buyfive)
		fmt.Printf("\n")
		fmt.Printf("------------------------- %v -----------------------------\n", ret.Retdata.Market.Shanghai.Name)
		fmt.Printf("current dot     : %v\n", ret.Retdata.Market.Shanghai.Curdot)
		fmt.Printf("current price   : %v\n", ret.Retdata.Market.Shanghai.Curprice)
		fmt.Printf("growth rate     : %v%%\n", ret.Retdata.Market.Shanghai.Rate)
		fmt.Printf("deal number     : %v\n", ret.Retdata.Market.Shanghai.Dealnumber)

		time.Sleep(1 * time.Second)
	}

	os.Exit(0)
}
