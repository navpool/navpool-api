package communityFund

import (
	"github.com/NavPool/navpool-api/config"
	"github.com/NavPool/navpool-api/navcoind"
)

var acceptedVotes = map[string]bool{
	"yes":    true,
	"no":     true,
	"remove": true,
}

func PostProposalVote(address string, hash string, vote string, signature string) (success bool, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	if config.Get().Signature {
		validSignature, err := nav.VerifyMessage(address, signature, address+hash+vote)
		if err != nil || validSignature == false {
			return false, ErrSignatureNotValid
		}
	}

	_, err = nav.GetProposal(hash)
	if err != nil {
		return false, ErrProposalNotValid
	}

	return nav.ProposalVote(address, hash, vote)
}

func PostPaymentRequestVote(address string, hash string, vote string, signature string) (success bool, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	if config.Get().Signature {
		validSignature, err := nav.VerifyMessage(address, signature, address+hash+vote)
		if err != nil || validSignature == false {
			return false, ErrSignatureNotValid
		}
	}

	_, err = nav.GetPaymentRequest(hash)
	if err != nil {
		return false, ErrPaymentRequestNotValid
	}

	return nav.PaymentRequestVote(address, hash, vote)
}
