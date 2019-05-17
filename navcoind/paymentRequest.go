package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
)

type PaymentRequest struct {
	Version             int    `json:"version"`
	Hash                string `json:"hash"`
	BlockHash           string `json:"blockHash"`
	Description         string `json:"description"`
	RequestedAmount     string `json:"requestedAmount"`
	VotesYes            int    `json:"votesYes"`
	VotesNo             int    `json:"votesNo"`
	VotingCycle         int    `json:"votingCycle"`
	Status              string `json:"status"`
	State               int    `json:"state"`
	StateChangedOnBlock string `json:"stateChangedOnBlock"`
}

func (nav *Navcoind) ListPaymentRequestVotes(hash string) (votes []Votes, err error) {
	response, err := nav.client.call("poolpaymentrequestvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return votes, err
}

func (nav *Navcoind) GetPaymentRequest(hash string) (paymentRequest PaymentRequest, err error) {
	response, err := nav.client.call("getpaymentrequest", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &paymentRequest)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return paymentRequest, err
}

func (nav *Navcoind) PaymentRequestVote(address string, hash string, vote string) (success bool, err error) {
	response, err := nav.client.call("poolpaymentrequestvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}
