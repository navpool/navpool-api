package navcoind

import (
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/helpers"
)

type Navcoind struct {
	client *rpcClient
}

func New() (*Navcoind, error) {
	network := config.Get().Networks[0]
	if config.Get().SelectedNetwork == "testnet" {
		network = config.Get().Networks[1]
	}

	rpcClient, err := newClient(network.Host, network.Port, network.Username, network.Password)
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
