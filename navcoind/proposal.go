package navcoind

import (
	"encoding/json"
)

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

type Votes struct {
	Proposals []string
}

func (nav *Navcoind) ListProposalVotes(hash string) (votes []Votes, err error) {
	response, err := nav.client.call("poolproposalvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &votes)

	return votes, err
}

func (nav *Navcoind) GetProposal(hash string) (proposal Proposal, err error) {
	response, err := nav.client.call("getproposal", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &proposal)

	return proposal, err
}

func (nav *Navcoind) ProposalVote(address string, hash string, vote string) (success bool, err error) {
	response, err := nav.client.call("poolproposalvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}
