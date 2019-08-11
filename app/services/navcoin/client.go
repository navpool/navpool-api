package navcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NavPool/navpool-api/app/helpers"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	serverAddr string
	user       string
	password   string
	httpClient *http.Client
}

type Payload struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
}

type Response struct {
	Id     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
	Err    interface{}     `json:"error"`
}

func NewClient(host string, port int, user, password string) *Client {
	return &Client{
		serverAddr: fmt.Sprintf("http://%s:%d", host, port),
		user:       user,
		password:   password,
		httpClient: &http.Client{
			Timeout: time.Second,
		},
	}
}

func (c *Client) call(method string, params interface{}) (rr Response, err error) {
	payload := Payload{method, params, time.Now().UnixNano(), "1.0"}
	payloadBuffer := &bytes.Buffer{}

	err = json.NewEncoder(payloadBuffer).Encode(payload)
	if err != nil {
		helpers.LogError(err)
		return
	}

	helpers.Debugf("Navcoind: Request(%s)", payloadBuffer)
	req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)
	if err != nil {
		helpers.LogError(err)
		return
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	if len(c.user) > 0 || len(c.password) > 0 {
		req.SetBasicAuth(c.user, c.password)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		helpers.LogError(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	helpers.Debugf("Navcoind: Response(%s)", data)
	if err != nil {
		helpers.LogError(err)
		return
	}

	err = json.Unmarshal(data, &rr)
	if err != nil {
		helpers.LogError(err)
	}

	return
}
