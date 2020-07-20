package navcoind

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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

func (n *Navcoind) GetStakingInfo() (stakingInfo *StakingInfo, err error) {
	response, err := n.client.call("getstakinginfo", nil)
	if err = HandleError(err, &response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &stakingInfo)
	if err != nil {
		log.WithError(err).Error("Failed to unmarshall staking info")
	}

	return
}
