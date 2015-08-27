package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type stockHttp struct {
	params  url.Values
	url     string
	headers map[string]string
}

func (req *stockHttp) SetParams(params map[string]string) {
	if params == nil {
		return
	}

	if req.params == nil {
		req.params = url.Values{}
	}

	for k, v := range params {
		req.params.Add(k, v)
	}
}

func (req *stockHttp) SetHeaders(headers map[string]string) {
	if headers == nil {
		return
	}

	if req.headers == nil {
		req.headers = map[string]string{}
	}

	for k, v := range headers {
		req.headers[k] = v
	}
}

func (req *stockHttp) SetUrl(url string) {
	req.url = url
}

func (req *stockHttp) Request(ret interface{}) error {
	var hdl *http.Request
	var resp *http.Response
	var err error

	if hdl, err = http.NewRequest("GET", req.url+"?"+req.params.Encode(), nil); err != nil {
		return err
	}

	hdl.Form = req.params
	if err = hdl.ParseForm(); err != nil {
		return err
	}

	for k, v := range req.headers {
		hdl.Header.Add(k, v)
	}

	if resp, err = http.DefaultClient.Do(hdl); err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, ret); err != nil {
		return err
	}

	return nil
}
