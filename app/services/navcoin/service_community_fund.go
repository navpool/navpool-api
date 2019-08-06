package navcoin

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/app/helpers"
)

func (nav *Navcoin) GetProposal(hash string) (proposal Proposal, err error) {
	response, err := nav.Client.call("getproposal", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &proposal)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

func (nav *Navcoin) GetPaymentRequest(hash string) (paymentRequest PaymentRequest, err error) {
	response, err := nav.Client.call("getpaymentrequest", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &paymentRequest)
	if err != nil {
		helpers.LogError(err)
	}

	return paymentRequest, err
}

func (nav *Navcoin) ListProposalVotes(hash string) (votes []Votes, err error) {
	response, err := nav.Client.call("poolproposalvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		helpers.LogError(err)
	}

	return votes, err
}

func (nav *Navcoin) ListPaymentRequestVotes(hash string) (votes []Votes, err error) {
	response, err := nav.Client.call("poolpaymentrequestvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		helpers.LogError(err)
	}

	return votes, err
}

func (nav *Navcoin) ProposalVote(address string, hash string, vote string) (success bool, err error) {
	response, err := nav.Client.call("poolproposalvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}

func (nav *Navcoin) PaymentRequestVote(address string, hash string, vote string) (success bool, err error) {
	response, err := nav.Client.call("poolpaymentrequestvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}

type Proposal struct {
	Version          int     `json:"version"`
	Hash             string  `json:"hash"`
	BlockHash        string  `json:"blockHash"`
	Description      string  `json:"description"`
	RequestedNav     float64 `json:"requestNav"`
	NotPaidYet       string  `json:"notPaidYet"`
	UserPaidFee      string  `json:"userPaidFee"`
	PaymentAddress   string  `json:"paymentAddress"`
	ProposalDuration int     `json:"proposalDuration"`
	VotesYes         int     `json:"votesYes"`
	VotesNo          int     `json:"votesNo"`
	VotingCycle      int     `json:"votingCycle"`
	Status           string  `json:"status"`
	State            int     `json:"state"`
}

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

type Votes struct {
	Proposals []string
}

var AcceptedVotes = map[string]bool{
	"yes":    true,
	"no":     true,
	"remove": true,
}
