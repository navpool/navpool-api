package controllers

import (
	"errors"
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/helpers"
	"github.com/NavPool/navpool-api/navcoind"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommunityFundController struct{}

func (controller *CommunityFundController) GetProposalVotes(c *gin.Context) {
	votes, err := navcoind.CommunityFund{}.ListProposalVotes(c.Param("address"))

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

func (controller *CommunityFundController) GetPaymentRequestVotes(c *gin.Context) {
	votes, err := navcoind.CommunityFund{}.ListPaymentRequestVotes(c.Param("address"))

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

func (controller *CommunityFundController) PostProposalVote(c *gin.Context) {
	var vote vote

	err := c.BindJSON(&vote)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if !navcoind.AcceptedVotes[vote.Vote] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if config.Get().Signature {
		validSignature, err := navcoind.Signature{}.VerifySignature(vote.SpendingAddress, vote.Signature, vote.SpendingAddress+vote.Hash+vote.Vote)
		if err != nil || validSignature == false {
			helpers.HandleError(c, ErrSignatureNotValid, http.StatusBadRequest)
			return
		}
	}

	success, err := navcoind.CommunityFund{}.ProposalVote(vote.SpendingAddress, vote.Hash, vote.Vote)
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

func (controller *CommunityFundController) PostPaymentRequestVote(c *gin.Context) {
	var vote vote

	err := c.BindJSON(&vote)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	communityFundService := navcoind.CommunityFund{}
	signatureService := navcoind.Signature{}

	if !navcoind.AcceptedVotes[vote.Vote] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if config.Get().Signature {
		validSignature, err := signatureService.VerifySignature(vote.SpendingAddress, vote.Signature, vote.SpendingAddress+vote.Hash+vote.Vote)
		if err != nil || validSignature == false {
			helpers.HandleError(c, ErrSignatureNotValid, http.StatusBadRequest)
			return
		}
	}

	success, err := communityFundService.PaymentRequestVote(vote.SpendingAddress, vote.Hash, vote.Vote)
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
