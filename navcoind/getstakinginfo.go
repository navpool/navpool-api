package navcoind

import (
	"encoding/json"
	"github.com/NavPool/navpool-api/logger"
)

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

func (nav *Navcoind) GetStakingInfo() (stakingInfo StakingInfo, err error) {
	response, err := nav.client.call("getstakinginfo", nil)
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &stakingInfo)
	if err != nil {
		logger.LogError(err)
	}

	return
}
