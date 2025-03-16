package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}
}
