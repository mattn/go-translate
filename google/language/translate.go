package translate

import (
	"bytes"
	"errors"
	"net/url"
)
import "net/http"
import "net/http/httputil"
import "io/ioutil"
import "encoding/json"
import "net"
import "os"

func Translate(from, to, s string) (t string, err error) {
	var url_ = "http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&format=text&q=" + url.QueryEscape(s) + "&langpair=" + url.QueryEscape(from) + "%7C" + url.QueryEscape(to)
	var r *http.Response
	if proxy := os.Getenv("HTTP_PROXY"); len(proxy) > 0 {
		proxy_url, _ := url.Parse(proxy)
		tcp, _ := net.Dial("tcp", proxy_url.Host)
		conn := httputil.NewClientConn(tcp, nil)
		var req http.Request
		req.URL, _ = url.Parse(url_)
		req.Method = "GET"
		r, err = conn.Do(&req)
	} else {
		r, err = http.Get(url_)
	}
	if err == nil {
		defer r.Body.Close()
		if b, err := ioutil.ReadAll(r.Body); err == nil {
			var r interface{}
			if err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&r); err == nil {
				if r.(map[string]interface{})["responseStatus"].(float64) == 200 {
					return r.(map[string]interface{})["responseData"].(map[string]interface{})["translatedText"].(string), nil
				} else {
					err = errors.New(r.(map[string]interface{})["responseDetails"].(string))
				}
			}
		}
	}
	return "", err
}
