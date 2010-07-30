package translate

import "http"
import "os"
import "json"
import "io/ioutil"
import "bytes"

func Translate(from, to, s string) (t string, err os.Error) {
	if r, _, err := http.Get("http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&q=" + http.URLEscape(s) + "&langpair=" + http.URLEscape(from) + "%7C" + http.URLEscape(to)); err == nil {
		defer r.Body.Close()
		if b, err := ioutil.ReadAll(r.Body); err == nil {
			var r interface{};
			if err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&r); err == nil {
				return r.(map[string]interface{})["responseData"].(map[string]interface{})["translatedText"].(string), nil
			}
		}
	}
	return "", err;
}
