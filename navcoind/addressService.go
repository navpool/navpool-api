package navcoind

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/helpers"
)

type Address struct{}

func (a Address) GetPoolAddress(spendingAddress string) (poolAddress PoolAddress, err error) {
	nav, err := New()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("newpooladdress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &poolAddress)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

func (a Address) GetValidateAddress(spendingAddress string) (validateAddress ValidateAddress, err error) {
	nav, err := New()
	if err != nil {
		helpers.LogError(err)
		return
	}

	response, err := nav.client.call("validateaddress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &validateAddress)
	if err != nil {
		helpers.LogError(err)
		return
	}

	return
}

type PoolAddress struct {
	SpendingAddress    string `json:"spendingAddress"`
	StakingAddress     string `json:"stakingAddress"`
	ColdStakingAddress string `json:"coldStakingAddress"`
}

type ValidateAddress struct {
	Valid           bool   `json:"isvalid"`
	Address         string `json:"address"`
	StakingAddress  string `json:"stakingaddress"`
	SpendingAddress string `json:"spendingaddress"`
	ScriptPubKey    string `json:"scriptpubkey"`
	Mine            bool   `json:"ismine"`
	Stakeable       bool   `json:"isstakeable"`
	WatchOnly       bool   `json:"iswatchonly"`
	Script          bool   `json:"isscript"`
	ColdStaking     bool   `json:"iscoldstaking"`
	PubKey          string `json:"pubkey"`
	StakingPubKey   string `json:"stakingpubkey"`
	SpendingPubKey  string `json:"spendingpubkey"`
	Compressed      bool   `json:"iscompressed"`
	Account         string `json:"account"`
	HdKeyPath       string `json:"hdkeypath"`
	HdMasterKey     string `json:"hdmasterkey"`
}
