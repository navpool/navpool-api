package di

import (
	"github.com/NavPool/navpool-api/internal/navcoind"
	"github.com/NavPool/navpool-api/internal/service"
	"github.com/sarulabs/dingo/v3"
)

var Definitions = []dingo.Def{
	{
		Name: "navcoind",
		Build: func() (*navcoind.Factory, error) {
			return &navcoind.Factory{}, nil
		},
	},
	{
		Name: "address.service",
		Build: func(n *navcoind.Factory) (*service.AddressService, error) {
			return service.NewAddressService(n), nil
		},
	},
	{
		Name: "dao.service",
		Build: func(n *navcoind.Factory) (*service.DaoService, error) {
			return service.NewDaoService(n), nil
		},
	},
	{
		Name: "info.service",
		Build: func(n *navcoind.Factory) (*service.InfoService, error) {
			return service.NewInfoService(n), nil
		},
	},
}
