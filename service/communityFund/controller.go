package communityFund

import (
	"errors"
	"github.com/NavPool/navpool-api/error"
	"github.com/NavPool/navpool-api/service/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func (controller *Controller) PostVote(c *gin.Context) {
	vote := c.PostForm("vote")
	if !acceptedVotes[vote] {
		error.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	addressHash := c.PostForm("address")
	if !address.IsValid(addressHash) {
		error.HandleError(c, ErrAddressNotValid, http.StatusBadRequest)
		return
	}

	hash := c.PostForm("hash")
	signature := c.PostForm("signature")

	voteType := c.Param("type")

	if voteType != "proposal" && voteType != "payment-request" {
		error.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		return
	}

	if voteType == "proposal" {
		success, err := PostProposalVote(addressHash, hash, vote, signature)
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
		success, err := PostPaymentRequestVote(addressHash, hash, vote, signature)
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
