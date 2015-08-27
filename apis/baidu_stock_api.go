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
	Errnum  int    `json:"errNum,omitempty"`
	Errmsg  string `json:"errMsg,omitempty"`
	Retdata struct {
		Stockinfo []struct {
			Name          string  `json:"name,omitempty"`
			Code          string  `json:"code,omitempty"`
			Date          string  `json:"date,omitempty"`
			Openningprice float64 `json:"openningPrice,omitempty"`
			Closingprice  float64 `json:"closingPrice,omitempty"`
			Hprice        float64 `json:"hPrice,omitempty"`
			Lprice        float64 `json:"lPrice,omitempty"`
			Currentprice  float64 `json:"currentPrice,omitempty"`
			Growth        float64 `json:"growth,omitempty"`
			Growthpercent float64 `json:"growthPercent,omitempty"`
			Dealnumber    float64 `json:"dealnumber,omitempty"`
			Turnover      float64 `json:"turnover,omitempty"`
			Five2Hprice   float64 `json:"52hPrice,omitempty"`
			Five2Lprice   float64 `json:"52lPrice,omitempty"`
		} `json:"stockinfo,omitempty"`
		Market struct {
			Shanghai struct {
				Name       string  `json:"name,omitempty"`
				Curdot     float64 `json:"curdot,omitempty"`
				Curprice   float64 `json:"curprice,omitempty"`
				Rate       float64 `json:"rate,omitempty"`
				Dealnumber float64 `json:"dealnumber,omitempty"`
				Turnover   float64 `json:"turnover,omitempty"`
			} `json:"shanghai,omitempty"`
			Shenzhen struct {
				Name       string  `json:"name,omitempty"`
				Curdot     float64 `json:"curdot,omitempty"`
				Curprice   float64 `json:"curprice,omitempty"`
				Rate       float64 `json:"rate,omitempty"`
				Dealnumber float64 `json:"dealnumber,omitempty"`
				Turnover   float64 `json:"turnover,omitempty"`
			} `json:"shenzhen,omitempty"`
			Dji struct {
				Name     string  `json:"name,omitempty"`
				Date     string  `json:"date,omitempty"`
				Curdot   float64 `json:"curdot,omitempty"`
				Rate     float64 `json:"rate,omitempty"`
				Growth   float64 `json:"growth,omitempty"`
				Startdot float64 `json:"startdot,omitempty"`
				Closedot float64 `json:"closedot,omitempty"`
				Hdot     float64 `json:"hdot,omitempty"`
				Ldot     float64 `json:"ldot,omitempty"`
				Turnover int     `json:"turnover,omitempty"`
			} `json:"DJI,omitempty"`
			Ixic struct {
				Name     string  `json:"name,omitempty"`
				Date     string  `json:"date,omitempty"`
				Curdot   float64 `json:"curdot,omitempty"`
				Rate     float64 `json:"rate,omitempty"`
				Growth   float64 `json:"growth,omitempty"`
				Startdot float64 `json:"startdot,omitempty"`
				Closedot float64 `json:"closedot,omitempty"`
				Hdot     float64 `json:"hdot,omitempty"`
				Ldot     float64 `json:"ldot,omitempty"`
				Turnover int     `json:"turnover,omitempty"`
			} `json:"IXIC,omitempty"`
			Inx struct {
				Name     string  `json:"name,omitempty"`
				Date     string  `json:"date,omitempty"`
				Curdot   float64 `json:"curdot,omitempty"`
				Rate     float64 `json:"rate,omitempty"`
				Growth   float64 `json:"growth,omitempty"`
				Startdot float64 `json:"startdot,omitempty"`
				Closedot float64 `json:"closedot,omitempty"`
				Hdot     float64 `json:"hdot,omitempty"`
				Ldot     float64 `json:"ldot,omitempty"`
				Turnover int     `json:"turnover,omitempty"`
			} `json:"INX,omitempty"`
			Hsi struct {
				Name      string  `json:"name,omitempty"`
				Date      string  `json:"date,omitempty"`
				Curdot    float64 `json:"curdot,omitempty"`
				Rate      float64 `json:"rate,omitempty"`
				Growth    float64 `json:"growth,omitempty"`
				Startdot  float64 `json:"startdot,omitempty"`
				Closedot  float64 `json:"closedot,omitempty"`
				Hdot      float64 `json:"hdot,omitempty"`
				Ldot      float64 `json:"ldot,omitempty"`
				Turnover  float64 `json:"turnover,omitempty"`
				Five2Hdot float64 `json:"52hdot,omitempty"`
				Five2Ldot float64 `json:"52ldot,omitempty"`
			} `json:"HSI,omitempty"`
		} `json:"market,omitempty"`
	} `json:"retData,omitempty"`
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
