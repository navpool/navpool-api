package service

import (
	"errors"
	"github.com/NavPool/navpool-api/internal/navcoind"
)

type AddressService struct {
	navcoind *navcoind.Factory
}

func NewAddressService(navcoind *navcoind.Factory) *AddressService {
	return &AddressService{navcoind}
}

func (s *AddressService) GetPoolAddress(spendingAddress string) (*navcoind.PoolAddress, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.GetPoolAddress(spendingAddress)
}

func (s *AddressService) IsValid(address string) (bool, error) {
	validateAddress, err := s.ValidateAddress(address)
	if err != nil {
		return false, err
	}

	return validateAddress.Valid, nil
}

func (s *AddressService) ValidateAddress(spendingAddress string) (*navcoind.ValidateAddress, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.GetValidateAddress(spendingAddress)
}

func (s *AddressService) VerifySignature(address, signature, message string) (bool, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return false, err
	}

	return n.VerifyMessage(address, signature, message)
}

var (
	ErrUnableToValidateAddress = errors.New("unable to validate address")
	ErrAddressNotValid         = errors.New("address not valid")
	ErrAddressIsColdStaking    = errors.New("address is a cold staking address")
	ErrAddressIsMine           = errors.New("address is owned by the pool")
	ErrPoolAddressNotAvailable = errors.New("pool address not available")
	ErrSignatureNotValid       = errors.New("signature not valid")
)
