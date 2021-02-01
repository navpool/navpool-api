package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
)

type Proposal struct {
	Version             int              `json:"version"`
	Hash                string           `json:"hash"`
	BlockHash           string           `json:"blockHash"`
	Description         string           `json:"description"`
	RequestedAmount     string           `json:"requestedAmount"`
	NotPaidYet          string           `json:"notPaidYet"`
	NotRequestedYet     string           `json:"notRequestedYet"`
	UserPaidFee         string           `json:"userPaidFee"`
	OwnerAddress        string           `json:"ownerAddress"`
	PaymentAddress      string           `json:"paymentAddress"`
	ProposalDuration    uint64           `json:"proposalDuration"`
	ExpiresOn           uint64           `json:"expiresOn"`
	VotesYes            uint             `json:"votesYes"`
	VotesAbs            uint             `json:"votesAbs"`
	VotesNo             uint             `json:"votesNo"`
	VotingCycle         uint             `json:"votingCycle"`
	Status              string           `json:"status"`
	State               uint             `json:"state"`
	StateChangedOnBlock string           `json:"stateChangedOnBlock,omitempty"`
	PaymentRequests     []PaymentRequest `json:"paymentRequests"`
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
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return votes, err
}

func (nav *Navcoind) GetProposal(hash string) (proposal Proposal, err error) {
	response, err := nav.client.call("getproposal", []interface{}{hash})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &proposal)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return proposal, err
}

func (nav *Navcoind) ProposalVote(address string, hash string, vote string) (success bool, err error) {
	response, err := nav.client.call("poolproposalvote", []interface{}{address, hash, vote})
	if err = HandleError(err, &response); err != nil {
		return false, err
	}

	return true, err
}
