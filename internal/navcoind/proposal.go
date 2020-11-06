package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
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
	Yes     []string `json:"yes"`
	No      []string `json:"no"`
	Abstain []string `json:"abs"`
}

func (n *Navcoind) ListProposalVotes(hash string) (votes *Votes, err error) {
	response, err := n.client.call("poolproposalvotelist", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	log.Info(string(response.Result))
	err = json.Unmarshal(response.Result, &votes)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return votes, err
}

func (n *Navcoind) GetProposal(hash string) (proposal Proposal, err error) {
	response, err := n.client.call("getproposal", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &proposal)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return proposal, err
}

func (n *Navcoind) ProposalVote(address string, hash string, vote string) (success bool, err error) {
	response, err := n.client.call("poolproposalvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}
