package navcoin

import (
	"encoding/json"
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	model_community_fund "github.com/NavPool/navpool-api/app/model/community_fund"
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

func (nav *Navcoin) SetVote(voteType model_community_fund.VoteType, address string, hash string, vote string) (bool, error) {
	var method string
	if voteType == model_community_fund.VoteTypeProposal {
		method = "poolproposalvote"
	} else if voteType == model_community_fund.VoteTypePaymentRequest {
		method = "poolpaymentrequestvote"
	} else {
		return false, errors.New("Invalid vote type")
	}

	response, err := nav.Client.call(method, []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, nil
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
