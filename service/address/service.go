package address

import (
	"github.com/NavPool/navpool-api/navcoind"
)

func GetPoolAddress(spendingAddress string) (poolAddress navcoind.PoolAddress, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	return nav.GetPoolAddress(spendingAddress)
}

func IsValid(address string) (valid bool) {
	validateAddress, err := ValidateAddress(address)
	if err != nil {
		return false
	}

	return validateAddress.Valid
}

func ValidateAddress(spendingAddress string) (validateAddress navcoind.ValidateAddress, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	validateAddress, err = nav.GetValidateAddress(spendingAddress)
	if err != nil {
		return
	}

	return validateAddress, err
}

func VerifySignature(address string, signature string, message string) (valid bool, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	return nav.VerifyMessage(address, signature, message)
}
