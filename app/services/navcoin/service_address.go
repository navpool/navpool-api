package navcoin

import (
	"encoding/json"
	"errors"
	"github.com/NavPool/navpool-api/app/helpers"
	uuid "github.com/satori/go.uuid"
)

func (nav *Navcoin) CreatePoolAddress(spendingAddress string, uuid uuid.UUID) (poolAddress PoolAddress, err error) {
	response, err := nav.Client.call("createpooladdress", []interface{}{spendingAddress, uuid.String()})
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &poolAddress)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

func (nav *Navcoin) ValidateAddress(spendingAddress string) (validateAddress ValidateAddress, err error) {
	response, err := nav.Client.call("validateaddress", []interface{}{spendingAddress})
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
	Uuid               uuid.UUID `json:"uuid"`
	SpendingAddress    string    `json:"spendingaaddress"`
	StakingAddress     string    `json:"stakingaddress"`
	ColdStakingAddress string    `json:"coldstakingaddress"`
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

var (
	ErrUnableToCreateAddress = errors.New("Unable to create address")
)
