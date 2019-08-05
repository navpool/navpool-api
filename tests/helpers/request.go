package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetBody(req *http.Request) (body []byte) {
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		return body
	}

	return
}
