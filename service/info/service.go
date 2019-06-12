package info

import (
	"github.com/NavPool/navpool-api/logger"
	"github.com/NavPool/navpool-api/navcoind"
)

func GetInfo() (info navcoind.Info, err error) {
	nav, err := navcoind.New()
	if err != nil {
		return
	}

	return nav.GetInfo()
}

func GetStakingInfo() (stakingInfo navcoind.StakingInfo, err error) {
	nav, err := navcoind.New()
	if err != nil {
		logger.LogError(err)
		return
	}

	stakingInfo, err = nav.GetStakingInfo()
	if err != nil {
		logger.LogError(err)
	}

	return
}
