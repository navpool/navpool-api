package navcoind

import (
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/session"
)

type Navcoind struct {
	client *rpcClient
}

func NewNavcoind() (*Navcoind, error) {
	network := config.Get().Networks[0]
	if session.Network == "testnet" {
		network = config.Get().Networks[1]
	}

	rpcClient, err := NewRpcClient(network.Host, network.Port, network.Username, network.Password)
	if err != nil {
		helpers.LogError(err)
		return nil, err
	}

	return &Navcoind{rpcClient}, nil
}

func HandleError(err error, r *rpcResponse) error {
	if err != nil {
		helpers.LogError(err)
		return err
	}

	if r.Err != nil {
		helpers.LogError(err)
		rr := r.Err.(map[string]interface{})
		return errors.New(fmt.Sprintf("(%v) %s", rr["code"].(float64), rr["message"].(string)))
	}

	return nil
}
