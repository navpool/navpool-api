package navcoin

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/app/helpers"
)

func (nav *Navcoin) VerifyMessage(address string, signature string, message string) (valid bool, err error) {
	response, err := nav.Client.call("verifymessage", []interface{}{address, signature, message})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	err = json.Unmarshal(response.Result, &valid)
	if err != nil {
		helpers.LogError(err)
	}

	return valid, err
}
