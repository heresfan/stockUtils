package apis

import (
	"fmt"
)

const (
	BAIDU_STOCK_API    = "http://apis.baidu.com/apistore/stockservice/stock"
	BAIDU_STOCK_APIKEY = "8a6599d69fca87efbc29a63300d0c9c2"
)

const (
	REQ_TYPE_MULTI  = "1"
	REQ_TYPE_SINGLE = "2"
)

type BaiduStoctInfo struct {
	Errnum  int    `json:"errNum"`
	Errmsg  string `json:"errMsg"`
	Retdata struct {
		Stockinfo []struct {
			Name             string  `json:"name"`
			Code             string  `json:"code"`
			Date             string  `json:"date"`
			Time             string  `json:"time"`
			Openningprice    float64 `json:"OpenningPrice"`
			Closingprice     float64 `json:"closingPrice"`
			Currentprice     float64 `json:"currentPrice"`
			Hprice           float64 `json:"hPrice"`
			Lprice           float64 `json:"lPrice"`
			Competitiveprice float64 `json:"competitivePrice"`
			Auctionprice     float64 `json:"auctionPrice"`
			Totalnumber      int     `json:"totalNumber"`
			Turnover         int     `json:"turnover"`
			Increase         float64 `json:"increase"`
			Buyone           int     `json:"buyOne"`
			Buyoneprice      float64 `json:"buyOnePrice"`
			Buytwo           int     `json:"buyTwo"`
			Buytwoprice      float64 `json:"buyTwoPrice"`
			Buythree         int     `json:"buyThree"`
			Buythreeprice    float64 `json:"buyThreePrice"`
			Buyfour          int     `json:"buyFour"`
			Buyfourprice     float64 `json:"buyFourPrice"`
			Buyfive          int     `json:"buyFive"`
			Buyfiveprice     float64 `json:"buyFivePrice"`
			Sellone          int     `json:"sellOne"`
			Selloneprice     float64 `json:"sellOnePrice"`
			Selltwo          int     `json:"sellTwo"`
			Selltwoprice     float64 `json:"sellTwoPrice"`
			Sellthree        int     `json:"sellThree"`
			Sellthreeprice   float64 `json:"sellThreePrice"`
			Sellfour         int     `json:"sellFour"`
			Sellfourprice    float64 `json:"sellFourPrice"`
			Sellfive         int     `json:"sellFive"`
			Sellfiveprice    float64 `json:"sellFivePrice"`
			Minurl           string  `json:"minurl"`
			Dayurl           string  `json:"dayurl"`
			Weekurl          string  `json:"weekurl"`
			Monthurl         string  `json:"monthurl"`
		} `json:"stockinfo"`
		Market struct {
			Shanghai struct {
				Name       string  `json:"name"`
				Curdot     float64 `json:"curdot"`
				Curprice   float64 `json:"curprice"`
				Rate       float64 `json:"rate"`
				Dealnumber int     `json:"dealnumber"`
				Turnover   int     `json:"turnover"`
			} `json:"shanghai"`
			Shenzhen struct {
				Name       string  `json:"name"`
				Curdot     float64 `json:"curdot"`
				Curprice   float64 `json:"curprice"`
				Rate       float64 `json:"rate"`
				Dealnumber int     `json:"dealnumber"`
				Turnover   int     `json:"turnover"`
			} `json:"shenzhen"`
			Dji struct {
				Name     string  `json:"name"`
				Date     string  `json:"date"`
				Curdot   float64 `json:"curdot"`
				Rate     float64 `json:"rate"`
				Growth   float64 `json:"growth"`
				Startdot float64 `json:"startdot"`
				Closedot float64 `json:"closedot"`
				Hdot     float64 `json:"hdot"`
				Ldot     float64 `json:"ldot"`
				Turnover int     `json:"turnover"`
			} `json:"DJI"`
			Ixic struct {
				Name     string  `json:"name"`
				Date     string  `json:"date"`
				Curdot   float64 `json:"curdot"`
				Rate     float64 `json:"rate"`
				Growth   float64 `json:"growth"`
				Startdot float64 `json:"startdot"`
				Closedot float64 `json:"closedot"`
				Hdot     float64 `json:"hdot"`
				Ldot     float64 `json:"ldot"`
				Turnover int64   `json:"turnover"`
			} `json:"IXIC"`
			Inx struct {
				Name     string  `json:"name"`
				Date     string  `json:"date"`
				Curdot   float64 `json:"curdot"`
				Rate     float64 `json:"rate"`
				Growth   float64 `json:"growth"`
				Startdot float64 `json:"startdot"`
				Closedot float64 `json:"closedot"`
				Hdot     float64 `json:"hdot"`
				Ldot     float64 `json:"ldot"`
				Turnover int     `json:"turnover"`
			} `json:"INX"`
			Hsi struct {
				Name      string  `json:"name"`
				Date      string  `json:"date"`
				Curdot    float64 `json:"curdot"`
				Rate      float64 `json:"rate"`
				Growth    float64 `json:"growth"`
				Startdot  float64 `json:"startdot"`
				Closedot  float64 `json:"closedot"`
				Hdot      float64 `json:"hdot"`
				Ldot      float64 `json:"ldot"`
				Turnover  int     `json:"turnover"`
				Five2Hdot float64 `json:"52hdot"`
				Five2Ldot float64 `json:"52ldot"`
			} `json:"HSI"`
		} `json:"market"`
	} `json:"retData"`
}

func BaiduStoctRequest(sid string, ret *BaiduStoctInfo) (*BaiduStoctInfo, error) {
	if ret == nil {
		ret = new(BaiduStoctInfo)
	}

	var hdl stockHttp

	hdl.SetParams(map[string]string{
		"stockid": sid,
		"list":    REQ_TYPE_MULTI,
	})

	hdl.SetUrl(BAIDU_STOCK_API)

	hdl.SetHeaders(map[string]string{
		"apikey": BAIDU_STOCK_APIKEY,
	})

	if err := hdl.Request(ret); err != nil {
		return nil, err
	}

	if ret.Errnum != 0 {
		return nil, fmt.Errorf("Ret err[no:%v, msg:%v]", ret.Errnum, ret.Errmsg)
	}
	return ret, nil
}
