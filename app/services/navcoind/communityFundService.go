package navcoind

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/app/helpers"
)

type CommunityFund struct{}

func (c CommunityFund) GetProposal(hash string) (proposal Proposal, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("getproposal", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &proposal)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

func (c CommunityFund) GetPaymentRequest(hash string) (paymentRequest PaymentRequest, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("getpaymentrequest", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &paymentRequest)
	if err != nil {
		helpers.LogError(err)
	}

	return paymentRequest, err
}

func (c CommunityFund) ListProposalVotes(hash string) (votes []Votes, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("poolproposalvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		helpers.LogError(err)
	}

	return votes, err
}

func (c CommunityFund) ListPaymentRequestVotes(hash string) (votes []Votes, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("poolpaymentrequestvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		helpers.LogError(err)
	}

	return votes, err
}

func (c CommunityFund) ProposalVote(address string, hash string, vote string) (success bool, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("poolproposalvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}

func (c CommunityFund) PaymentRequestVote(address string, hash string, vote string) (success bool, err error) {
	nav, err := NewNavcoind()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("poolpaymentrequestvote", []interface{}{address, hash, vote})
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
