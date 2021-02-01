package navcoind

import (
	"encoding/json"
	"github.com/getsentry/raven-go"
)

type Info struct {
	Version               int           `json:"version"`
	ProtocolVersion       int           `json:"protocolversion"`
	WalletVersion         int           `json:"walletversion"`
	Balance               float64       `json:"balance"`
	PrivateBalance        float64       `json:"private_balance"`
	PrivateBalancePending float64       `json:"private_balance_pending"`
	ColdStakingBalance    float64       `json:"coldstaking_balance"`
	NewMint               float64       `json:"newmint"`
	Stake                 float64       `json:"stake"`
	Blocks                int           `json:"blocks"`
	CommunityFund         CommunityFund `json:"communityfund"`
	PublicMoneySupply     float64       `json:"publicmoneysupply"`
	PrivateMoneySupply    float64       `json:"privatemoneysupply"`
	TimeOffset            int           `json:"timeoffset"`
	NtpTimeOffset         int           `json:"ntptimeoffset"`
	Connections           int           `json:"connections"`
	Proxy                 string        `json:"proxy"`
	TestNet               bool          `json:"testnet"`
	KeyPoolOldest         int           `json:"keypoololdest"`
	KeyPoolSize           int           `json:"keypoolsize"`
	UnlockedUntil         int           `json:"unlocked_until"`
	PayTxFee              float64       `json:"paytxfee"`
	RelayFee              float64       `json:"relayfee"`
	Errors                string        `json:"errors"`
}

type CommunityFund struct {
	Available float64 `json:"available"`
	Locked    float64 `json:"locked"`
}

func (nav *Navcoind) GetInfo() (info Info, err error) {
	response, err := nav.client.call("getinfo", nil)
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &info)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}

	return
}
