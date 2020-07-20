package service

import (
	"errors"
	"github.com/NavPool/navpool-api/internal/config"
	"github.com/NavPool/navpool-api/internal/navcoind"
)

type DaoService struct {
	navcoind *navcoind.Factory
}

func NewDaoService(navcoind *navcoind.Factory) *DaoService {
	return &DaoService{navcoind}
}

var acceptedVotes = map[string]bool{
	"yes":    true,
	"no":     true,
	"remove": true,
}

func (s *DaoService) IsVoteValid(vote string) bool {
	return acceptedVotes[vote]
}

func (s *DaoService) GetListProposalVotes(address string) ([]navcoind.Votes, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.ListProposalVotes(address)
}

func (s *DaoService) GetListPaymentRequestVotes(address string) ([]navcoind.Votes, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.ListPaymentRequestVotes(address)
}

func (s *DaoService) PostProposalVote(address, hash, vote, signature string) (bool, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return false, err
	}

	if config.Get().Signature {
		validSignature, err := n.VerifyMessage(address, signature, address+hash+vote)
		if err != nil || validSignature == false {
			return false, ErrSignatureNotValid
		}
	}

	_, err = n.GetProposal(hash)
	if err != nil {
		return false, ErrProposalNotValid
	}

	return n.ProposalVote(address, hash, vote)
}

func (s *DaoService) PostPaymentRequestVote(address, hash, vote, signature string) (bool, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return false, err
	}

	if config.Get().Signature {
		validSignature, err := n.VerifyMessage(address, signature, address+hash+vote)
		if err != nil || validSignature == false {
			return false, ErrSignatureNotValid
		}
	}

	_, err = n.GetPaymentRequest(hash)
	if err != nil {
		return false, ErrPaymentRequestNotValid
	}

	return n.PaymentRequestVote(address, hash, vote)
}

var (
	ErrVoteNotValid           = errors.New("vote not valid")
	ErrProposalNotValid       = errors.New("proposal not valid")
	ErrPaymentRequestNotValid = errors.New("payment request not valid")
	ErrUnableToCastVote       = errors.New("unable to cast vote")
)
