package resource

import (
	"errors"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/helpers"
	model_address "github.com/NavPool/navpool-api/app/model/address"
	model_community_fund "github.com/NavPool/navpool-api/app/model/community_fund"
	"github.com/NavPool/navpool-api/app/services/navcoin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommunityFundResource struct{}

func (r *CommunityFundResource) GetProposalVotes(c *gin.Context) {
	votes, err := navcoin.NewNavcoin().ListProposalVotes(c.Param("address"))

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
	votes, err := navcoin.NewNavcoin().ListPaymentRequestVotes(c.Param("address"))

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
	var voteDto voteDto

	err := c.BindJSON(&voteDto)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if !navcoin.AcceptedVotes[voteDto.Choice] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	_, err = model_address.AddressRepository().GetAddressByStakingAddress(container.Container.User, voteDto.StakingAddress)
	if err != nil {
		helpers.HandleError(c, ErrStakingAddressNotValid, http.StatusBadRequest)
		return
	}

	vote := model_community_fund.CommunityFundRepository().GetVote(
		container.Container.User,
		model_community_fund.VoteTypeProposal,
		voteDto.StakingAddress,
		voteDto.Hash,
		voteDto.Choice,
	)

	success, err := navcoin.NewNavcoin().SetVote(
		model_community_fund.VoteTypeProposal,
		vote.StakingAddress,
		vote.Hash,
		vote.Choice,
	)
	if err != nil {
		if err == ErrProposalNotValid {
			helpers.HandleError(c, err, http.StatusBadRequest)
		} else {
			helpers.HandleError(c, ErrUnableToCastVote, http.StatusInternalServerError)
		}
		return
	}

	model_community_fund.CommunityFundRepository().DB.Save(vote)

	c.JSON(200, success)
}

func (r *CommunityFundResource) PostPaymentRequestVote(c *gin.Context) {
	var voteDto voteDto

	err := c.BindJSON(&voteDto)
	if err != nil {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	if !navcoin.AcceptedVotes[voteDto.Choice] {
		helpers.HandleError(c, ErrVoteNotValid, http.StatusBadRequest)
		return
	}

	_, err = model_address.AddressRepository().GetAddressByStakingAddress(container.Container.User, voteDto.StakingAddress)
	if err != nil {
		helpers.HandleError(c, ErrStakingAddressNotValid, http.StatusBadRequest)
		return
	}

	vote := model_community_fund.CommunityFundRepository().GetVote(
		container.Container.User,
		model_community_fund.VoteTypeProposal,
		voteDto.StakingAddress,
		voteDto.Hash,
		voteDto.Choice,
	)

	success, err := navcoin.NewNavcoin().SetVote(
		model_community_fund.VoteTypePaymentRequest,
		vote.StakingAddress,
		vote.Hash,
		vote.Choice,
	)
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

type voteDto struct {
	Hash           string `json:"hash" binding:"required"`
	StakingAddress string `json:"staking_address" binding:"required"`
	Choice         string `json:"vote" binding:"required"`
}

var (
	ErrVoteNotValid           = errors.New("The vote is not valid")
	ErrUnableToCastVote       = errors.New("Unable to cast vote")
	ErrProposalNotValid       = errors.New("The proposal is not valid")
	ErrPaymentRequestNotValid = errors.New("The payment request is not valid")
	ErrStakingAddressNotValid = errors.New("The staking address is not valid")
)
