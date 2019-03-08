package navcoind

import (
	"encoding/json"
)

type PoolAddress struct {
	SpendingAddress    string `json:"spendingAddress"`
	StakingAddress     string `json:"stakingAddress"`
	ColdStakingAddress string `json:"coldStakingAddress"`
}

type ValidateAddress struct {
	IsValid bool `json:"isvalid"`
}

func (nav *Navcoind) GetPoolAddress(spendingAddress string) (poolAddress PoolAddress, err error) {
	response, err := nav.client.call("newpooladdress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &poolAddress)

	return
}

func (nav *Navcoind) GetValidateAddress(spendingAddress string) (validateAddress ValidateAddress, err error) {
	response, err := nav.client.call("validateaddress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &validateAddress)
	if err != nil {
		return
	}

	return validateAddress, err
}
