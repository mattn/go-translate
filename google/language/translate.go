package translate

import "bytes"
import "http"
import "io/ioutil"
import "json"
import "net"
import "os"

func Translate(from, to, s string) (t string, err os.Error) {
	var url = "http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&format=text&q=" + http.URLEscape(s) + "&langpair=" + http.URLEscape(from) + "%7C" + http.URLEscape(to)
	var r *http.Response;
	if proxy := os.Getenv("HTTP_PROXY"); len(proxy) > 0 {
		proxy_url, _ := http.ParseURL(proxy);
		tcp, _ := net.Dial("tcp", "", proxy_url.Host);
		conn := http.NewClientConn(tcp, nil);
		var req http.Request;
		req.URL, _ = http.ParseURL(url);
		req.Method = "GET";
		err = conn.Write(&req);
		r, err = conn.Read();
	} else {
		r, _, err = http.Get(url);
	}
	if err == nil {
		defer r.Body.Close()
		if b, err := ioutil.ReadAll(r.Body); err == nil {
			var r interface{};
			if err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&r); err == nil {
				if r.(map[string]interface{})["responseStatus"].(float64) == 200 {
					return r.(map[string]interface{})["responseData"].(map[string]interface{})["translatedText"].(string), nil
				} else {
					err = os.NewError ( r.(map[string]interface{})["responseDetails"].(string) )
				}
			}
		}
	}
	return "", err;
}
