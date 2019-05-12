package communityFund

import (
	"errors"
	"github.com/NavPool/navpool-api/error"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct{}

type Vote struct {
	Hash            string `json:"hash" binding:"required"`
	SpendingAddress string `json:"spending_address" binding:"required"`
	Vote            string `json:"vote" binding:"required"`
	Signature       string `json:"signature"`
}

func (controller *Controller) GetVotes(c *gin.Context) {
	voteType := c.Param("type")
	spendingAddress := c.Param("address")

	if voteType == "proposal" {
		votes, err := GetListProposalVotes(spendingAddress)
		if err != nil {
			if err == ErrProposalNotValid {
				error.HandleError(c, ErrProposalNotValid, http.StatusBadRequest)
			} else {
				error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, votes)
	}

	if voteType == "payment-request" {
		votes, err := GetListPaymentRequestVotes(spendingAddress)
		if err != nil {
			if err == ErrProposalNotValid {
				error.HandleError(c, ErrProposalNotValid, http.StatusBadRequest)
			} else {
				error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, votes)
	}
}

func (controller *Controller) PostVote(c *gin.Context) {
	var vote Vote
	c.BindJSON(&vote)

	log.Printf("Address: %s, Hash: %s, Vote: %s", vote.SpendingAddress, vote.Hash, vote.Vote)
	if !acceptedVotes[vote.Vote] {
		error.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	voteType := c.Param("type")
	if voteType != "proposal" && voteType != "payment-request" {
		error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		return
	}

	if voteType == "proposal" {
		success, err := PostProposalVote(vote.SpendingAddress, vote.Hash, vote.Vote, vote.Signature)
		if err != nil {
			if err == ErrProposalNotValid {
				error.HandleError(c, ErrProposalNotValid, http.StatusBadRequest)
			} else {
				error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, success)
	}

	if voteType == "payment-request" {
		success, err := PostPaymentRequestVote(vote.SpendingAddress, vote.Hash, vote.Vote, vote.Signature)
		if err != nil {
			if err == ErrPaymentRequestNotValid {
				error.HandleError(c, ErrPaymentRequestNotValid, http.StatusBadRequest)
			} else {
				error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, success)
	}
}

var (
	ErrAddressNotValid        = errors.New("address not valid")
	ErrVoteNotValid           = errors.New("vote not valid")
	ErrProposalNotValid       = errors.New("proposal not valid")
	ErrPaymentRequestNotValid = errors.New("payment request not valid")
	ErrSignatureNotValid      = errors.New("signature not valid")
	ErrUnableToCastVote       = errors.New("unable to cast vote")
)
