package service

import (
	"github.com/NavPool/navpool-api/internal/navcoind"
)

type InfoService struct {
	navcoind *navcoind.Factory
}

func NewInfoService(navcoind *navcoind.Factory) *InfoService {
	return &InfoService{navcoind}
}

func (s *InfoService) GetInfo() (*navcoind.Info, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.GetInfo()
}

func (s *InfoService) GetStakingInfo() (*navcoind.StakingInfo, error) {
	n, err := s.navcoind.Connect()
	if err != nil {
		return nil, err
	}

	return n.GetStakingInfo()
}
