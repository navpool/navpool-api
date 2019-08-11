package model_community_fund

import (
	"errors"
	"github.com/NavPool/navpool-api/app/container"
	"github.com/NavPool/navpool-api/app/database"
	"github.com/NavPool/navpool-api/app/helpers"
	model_user "github.com/NavPool/navpool-api/app/model/user"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func CommunityFundRepository() *repo {
	return &repo{
		DB: database.GetConnection(),
	}
}

func (r *repo) GetVote(user model_user.User, voteType VoteType, stakingAddress string, hash string, choice string) Vote {
	var vote Vote

	r.DB.Where(Vote{UserID: user.ID, Type: voteType, StakingAddress: stakingAddress, Hash: hash}).
		Assign(Vote{Choice: choice, Committed: false}).
		FirstOrInit(&vote)

	return vote
}

func (r *repo) GetProposalVotes(user model_user.User) ([]Vote, error) {
	var votes = make([]Vote, 0)
	err := r.DB.Where(&Vote{UserID: user.ID, Type: VoteTypeProposal}).Find(&votes).Error
	if err != nil {
		helpers.LogError(err)
		return nil, ErrorUnableToGetProposalVotes
	}

	return votes, nil
}

func (r *repo) GetPaymentRequestVotes(user model_user.User) ([]Vote, error) {
	var votes = make([]Vote, 0)
	err := r.DB.Where(&Vote{UserID: user.ID, Type: VoteTypePaymentRequest}).Find(&votes).Error
	if err != nil {
		helpers.LogError(err)
		return nil, ErrorUnableToGetPaymentRequestVotes
	}

	return votes, nil
}

func UpdateProposalVote(vote Vote, user model_user.User) error {

	votes, err := GetProposalVotes(user)
	if err != nil {
		logger.LogError(err)
		return err
	}

	tx := db.Begin()

	modifiedVotes := make([]model.Vote, 0)
	for _, voteDto := range voteDtos {
		vote, err := matchedVote(voteDto.Hash, model.VoteTypeProposal, votes)
		if err == nil {
			err = tx.Model(&vote).Updates(model.Vote{Choice: voteDto.Choice, Committed: false}).Error
		} else {
			logger.LogError(err)
			newVote := &model.Vote{
				UserID:    user.ID,
				Type:      model.VoteTypeProposal,
				Hash:      voteDto.Hash,
				Choice:    voteDto.Choice,
				Committed: false,
			}
			err = tx.Create(newVote).Error
			vote = *newVote
		}

		if err != nil {
			logger.LogError(err)
			tx.Rollback()
			return err
		}
		modifiedVotes = append(modifiedVotes, vote)
	}

	err = tx.Commit().Error
	if err != nil {
		logger.LogError(err)
		return
	}
	err = updatePoolVotes(modifiedVotes, user)
	if err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

var (
	ErrorUnableToGetProposalVotes       = errors.New("Unable to retrieve proposal votes")
	ErrorUnableToGetPaymentRequestVotes = errors.New("Unable to retrieve payment request votes")
	ErrorUnableToMatchVote              = errors.New("Unable to match vote")
)
