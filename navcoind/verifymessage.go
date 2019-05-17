package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
)

func (nav *Navcoind) VerifyMessage(address string, signature string, message string) (valid bool, err error) {
	response, err := nav.client.call("verifymessage", []interface{}{address, signature, message})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	err = json.Unmarshal(response.Result, &valid)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return valid, err
}
