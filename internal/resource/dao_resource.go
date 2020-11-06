package resource

import (
	"github.com/NavPool/navpool-api/internal/framework/error"
	"github.com/NavPool/navpool-api/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type DaoResource struct {
	service *service.DaoService
}

func NewDaoResource(service *service.DaoService) *DaoResource {
	return &DaoResource{service}
}

type Vote struct {
	Hash            string `json:"hash" binding:"required"`
	SpendingAddress string `json:"spending_address" binding:"required"`
	Vote            string `json:"vote" binding:"required"`
	Signature       string `json:"signature"`
}

func (r *DaoResource) GetVotes(c *gin.Context) {
	voteType := c.Param("type")
	spendingAddress := c.Param("address")

	if voteType == "proposal" {
		votes, err := r.service.GetListProposalVotes(spendingAddress)
		if err != nil {
			if err == service.ErrProposalNotValid {
				error.HandleError(c, err, http.StatusBadRequest)
			} else {
				error.HandleError(c, service.ErrUnableToRetrieveVotes, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, votes)
	}

	if voteType == "payment-request" {
		votes, err := r.service.GetListPaymentRequestVotes(spendingAddress)
		if err != nil {
			if err == service.ErrProposalNotValid {
				error.HandleError(c, err, http.StatusBadRequest)
			} else {
				error.HandleError(c, service.ErrUnableToRetrieveVotes, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, votes)
	}
}

func (r *DaoResource) PostVote(c *gin.Context) {
	var vote Vote
	c.BindJSON(&vote)

	log.WithFields(log.Fields{
		"address": vote.SpendingAddress,
		"hash":    vote.Hash,
		"vote":    vote.Vote,
	}).Info("PostVote")
	if !r.service.IsVoteValid(vote.Vote) {
		error.HandleError(c, service.ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	voteType := c.Param("type")
	if voteType != "proposal" && voteType != "payment-request" {
		error.HandleError(c, service.ErrUnableToCastVote, http.StatusInternalServerError)
		return
	}

	if voteType == "proposal" {
		success, err := r.service.PostProposalVote(vote.SpendingAddress, vote.Hash, vote.Vote, vote.Signature)
		if err != nil {
			if err == service.ErrProposalNotValid {
				error.HandleError(c, err, http.StatusBadRequest)
				return
			}

			error.HandleError(c, service.ErrUnableToCastVote, http.StatusInternalServerError)
			return
		}

		c.JSON(200, success)
	}

	if voteType == "payment-request" {
		success, err := r.service.PostPaymentRequestVote(vote.SpendingAddress, vote.Hash, vote.Vote, vote.Signature)
		if err != nil {
			if err == service.ErrPaymentRequestNotValid {
				error.HandleError(c, err, http.StatusBadRequest)
			} else {
				error.HandleError(c, service.ErrUnableToCastVote, http.StatusInternalServerError)
			}
			return
		}

		c.JSON(200, success)
	}
}
