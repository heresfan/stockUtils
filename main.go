package main

import (
	"./apis"
	"fmt"
	"os"
)

func main() {
	var sid = "sz000048"
	var ret apis.BaiduStoctInfo

	if _, err := apis.BaiduStoctRequest(sid, &ret); err != nil {
		fmt.Printf("Failed err[%v]\n", err)
		os.Exit(1)
	}

	fmt.Printf("name:%v\n", ret.Retdata.Stockinfo[0].Name)

	os.Exit(0)
}
