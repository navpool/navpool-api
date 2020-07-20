package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
)

type PoolAddress struct {
	SpendingAddress    string `json:"spendingAddress"`
	StakingAddress     string `json:"stakingAddress"`
	ColdStakingAddress string `json:"coldStakingAddress"`
}

type ValidateAddress struct {
	Valid           bool   `json:"isvalid"`
	Address         string `json:"address"`
	ScriptPubKey    string `json:"scriptpubkey"`
	StakingAddress  string `json:"stakingaddress"`
	SpendingAddress string `json:"spendingaddress"`
	VotingAddress   string `json:"votingaddress"`
	Mine            bool   `json:"ismine"`
	Stakeable       bool   `json:"isstakeable"`
	WatchOnly       bool   `json:"iswatchonly"`
	Script          bool   `json:"isscript"`
	ColdStaking     bool   `json:"iscoldstaking"`
	PubKey          string `json:"pubkey"`
	Compressed      bool   `json:"iscompressed"`
	Account         string `json:"account"`
	HdKeyPath       string `json:"hdkeypath"`
	HdMasterKey     string `json:"hdmasterkey"`
}

func (n *Navcoind) GetPoolAddress(spendingAddress string) (poolAddress *PoolAddress, err error) {
	response, err := n.client.call("newpooladdress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &poolAddress)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return
}

func (n *Navcoind) GetValidateAddress(spendingAddress string) (validateAddress *ValidateAddress, err error) {
	response, err := n.client.call("validateaddress", []interface{}{spendingAddress})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &validateAddress)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}

	return validateAddress, err
}
