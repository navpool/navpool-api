package address

import (
	"github.com/NavPool/navpool-api/navcoind"
)

func GetPoolAddress(spendingAddress string) (poolAddress navcoind.PoolAddress, err error) {
	nav, err := navcoind.New()

	return nav.GetPoolAddress(spendingAddress)
}

func ValidateAddress(spendingAddress string) (valid bool) {
	nav, err := navcoind.New()
	if err != nil {
		return false
	}

	validateAddress, err := nav.GetValidateAddress(spendingAddress)
	if err != nil {
		return false
	}

	return validateAddress.IsValid
}
