package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/NavPool/navpool-api/app/services/navcoin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommunityFundResource struct{}

func (r *CommunityFundResource) GetProposalVotes(c *gin.Context) {
	votes, err := navcoin.NewNavcoin(nil).ListProposalVotes(c.Param("address"))

	if err != nil {
		if err == ErrProposalNotValid {
			helpers.HandleError(c, err, http.StatusBadRequest)
		} else {
			helpers.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(200, votes)
}

func (r *CommunityFundResource) GetPaymentRequestVotes(c *gin.Context) {
	votes, err := navcoin.NewNavcoin(nil).ListPaymentRequestVotes(c.Param("address"))

	if err != nil {
		if err == ErrProposalNotValid {
			helpers.HandleError(c, err, http.StatusBadRequest)
		} else {
			helpers.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(200, votes)
}

func (r *CommunityFundResource) PostProposalVote(c *gin.Context) {
	var vote vote

	err := c.BindJSON(&vote)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if !navcoin.AcceptedVotes[vote.Vote] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	nav := navcoin.NewNavcoin(nil)

	if config.Get().Signature {
		validSignature, err := nav.VerifyMessage(vote.SpendingAddress, vote.Signature, vote.SpendingAddress+vote.Hash+vote.Vote)
		if err != nil || validSignature == false {
			helpers.HandleError(c, err, http.StatusBadRequest)
			return
		}
	}

	success, err := nav.ProposalVote(vote.SpendingAddress, vote.Hash, vote.Vote)
	if err != nil {
		if err == ErrProposalNotValid {
			helpers.HandleError(c, err, http.StatusBadRequest)
		} else {
			helpers.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(200, success)
}

func (r *CommunityFundResource) PostPaymentRequestVote(c *gin.Context) {
	var vote vote

	err := c.BindJSON(&vote)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if !navcoin.AcceptedVotes[vote.Vote] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	nav := navcoin.NewNavcoin(nil)

	if config.Get().Signature {
		validSignature, err := nav.VerifyMessage(vote.SpendingAddress, vote.Signature, vote.SpendingAddress+vote.Hash+vote.Vote)
		if err != nil || validSignature == false {
			helpers.HandleError(c, err, http.StatusBadRequest)
			return
		}
	}

	success, err := nav.PaymentRequestVote(vote.SpendingAddress, vote.Hash, vote.Vote)
	if err != nil {
		if err == ErrPaymentRequestNotValid {
			helpers.HandleError(c, err, http.StatusBadRequest)
		} else {
			helpers.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(200, success)
}

type vote struct {
	Hash            string `json:"hash" binding:"required"`
	SpendingAddress string `json:"spending_address" binding:"required"`
	Vote            string `json:"vote" binding:"required"`
	Signature       string `json:"signature"`
}

var (
	ErrVoteNotValid           = errors.New("vote not valid")
	ErrUnableToCastVote       = errors.New("unable to cast vote")
	ErrProposalNotValid       = errors.New("proposal not valid")
	ErrPaymentRequestNotValid = errors.New("payment request not valid")
)
