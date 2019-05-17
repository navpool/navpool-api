package navcoind

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type rpcClient struct {
	serverAddr string
	user       string
	password   string
	httpClient *http.Client
}

type rpcRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
}

type rpcResponse struct {
	Id     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
	Err    interface{}     `json:"error"`
}

func newClient(host string, port int, user, password string) (c *rpcClient, err error) {
	if len(host) == 0 {
		err = errors.New("bad call missing argument host")
		return
	}

	var httpClient *http.Client
	httpClient = &http.Client{}

	c = &rpcClient{serverAddr: fmt.Sprintf("http://%s:%d", host, port), user: user, password: password, httpClient: httpClient}

	return
}

func (c *rpcClient) call(method string, params interface{}) (rr rpcResponse, err error) {
	rpcR := rpcRequest{method, params, time.Now().UnixNano(), "1.0"}
	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	err = jsonEncoder.Encode(rpcR)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}

	req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)
	log.Printf("Navcoind: Request(%s)", payloadBuffer)

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	if len(c.user) > 0 || len(c.password) > 0 {
		req.SetBasicAuth(c.user, c.password)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	log.Printf("Navcoind: Response(%s)", data)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}

	err = json.Unmarshal(data, &rr)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return
}
