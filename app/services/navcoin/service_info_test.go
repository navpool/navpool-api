package navcoin

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var payload Payload
		if err := json.Unmarshal(body, &payload); err != nil {
			panic(err)
		}
		assert.Equal(t, payload.Method, "getinfo")
		assert.Nil(t, payload.Params)

		w.Write([]byte(okResponse))
	})

	httpClient, teardown := NewTestClient(h)
	defer teardown()

	nav := NewNavcoin(&Client{serverAddr: "http://localhost", httpClient: httpClient})

	info, err := nav.GetInfo()
	assert.Nil(t, err)
	assert.Equal(t, 4060000, info.Version)
	assert.Equal(t, 70020, info.ProtocolVersion)
	assert.Equal(t, 130000, info.WalletVersion)
}

const (
	okResponse = `{"result": {
	  "version": 4060000,
	  "protocolversion": 70020,
	  "walletversion": 130000,
	  "balance": 0.00000000,
	  "coldstaking_balance": 1418792.54619298,
	  "newmint": 0.00000000,
	  "stake": 0.00000000,
	  "blocks": 3366602,
	  "communityfund": {
		"available": 30228.49325096,
		"locked": 146954.69000000
	  },
	  "timeoffset": 0,
	  "ntptimeoffset": -1,
	  "connections": 8,
	  "proxy": "",
	  "testnet": false,
	  "keypoololdest": 1560279002,
	  "keypoolsize": 101,
	  "unlocked_until": 1575121481,
	  "paytxfee": 0.00100000,
	  "relayfee": 0.00010000,
	  "errors": ""
	}, "error":null, "id":1565122734827902000}`
)
