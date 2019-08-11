package navcoin

import (
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
)

type Navcoin struct {
	Client *Client
}

func NewNavcoin() *Navcoin {
	client := NewClient(
		container.Container.Network.Host,
		container.Container.Network.Port,
		container.Container.Network.Username,
		container.Container.Network.Password,
	)

	return &Navcoin{client}
}

func HandleError(err error, r *Response) error {
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
