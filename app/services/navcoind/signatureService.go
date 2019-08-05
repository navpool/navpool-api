package navcoind

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/app/helpers"
)

type Signature struct{}

func (s Signature) VerifySignature(address string, signature string, message string) (valid bool, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("verifymessage", []interface{}{address, signature, message})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	err = json.Unmarshal(response.Result, &valid)
	if err != nil {
		helpers.LogError(err)
	}

	return valid, err
}
