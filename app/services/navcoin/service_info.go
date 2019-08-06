package navcoin

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/app/helpers"
)

func (nav *Navcoin) GetInfo() (info Info, err error) {
	response, err := nav.Client.call("getinfo", nil)
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &info)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

func (nav *Navcoin) GetStakingInfo() (stakingInfo StakingInfo, err error) {
	response, err := nav.Client.call("getstakinginfo", nil)
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &stakingInfo)
	if err != nil {
		helpers.LogError(err)
	}

	return
}

type Info struct {
	Version            int           `json:"version"`
	ProtocolVersion    int           `json:"protocolversion"`
	WalletVersion      int           `json:"walletversion"`
	Balance            float64       `json:"balance"`
	ColdStakingBalance float64       `json:"coldstaking_balance"`
	NewMint            float64       `json:"newmint"`
	Stake              float64       `json:"stake"`
	Blocks             int           `json:"blocks"`
	CommunityFund      communityFund `json:"communityfund"`
	TimeOffset         int           `json:"timeoffset"`
	NtpTimeOffset      int           `json:"ntptimeoffset"`
	Connections        int           `json:"connections"`
	Proxy              string        `json:"proxy"`
	TestNet            bool          `json:"testnet"`
	KeyPoolOldest      int           `json:"keypoololdest"`
	KeyPoolSize        int           `json:"keypoolsize"`
	UnlockedUntil      int           `json:"unlocked_until"`
	PayTxFee           float64       `json:"paytxfee"`
	RelayFee           float64       `json:"relayfee"`
	Errors             string        `json:"errors"`
}

type communityFund struct {
	Available float64 `json:"available"`
	Locked    float64 `json:"locked"`
}

type StakingInfo struct {
	Enabled          bool    `json:"enabled"`
	Staking          bool    `json:"staking"`
	Errors           string  `json:"errors"`
	CurrentBlockSize int     `json:"currentblocksize"`
	CurrentBlockTx   int     `json:"currentblocktx"`
	Difficulty       float64 `json:"difficulty"`
	SearchInterval   int     `json:"search-interval"`
	Weight           int     `json:"weight"`
	NetStakeWeight   int     `json:"netstakeweight"`
	ExpectedTime     int     `json:"expectedtime"`
}
