package navcoind

import (
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/internal/config"
	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
)

type Factory struct{}

func (f *Factory) Connect() (*Navcoind, error) {
	network, err := config.Get().ActiveNetwork()
	if err != nil {
		log.WithError(err).Error("Failed to select network")
		return nil, err
	}

	log.Info("Navcoind network: ", network.Name)

	return newNavcoind(network.Host, network.Port, network.Username, network.Password), nil
}

type Navcoind struct {
	client *rpcClient
}

func newNavcoind(host string, port int, user string, password string) *Navcoind {
	return &Navcoind{
		newClient(host, port, user, password),
	}
}

func HandleError(err error, r *rpcResponse) error {
	if err != nil {
		log.WithError(err).Error("Navcoind: HandleError")
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	if r.Err != nil {
		raven.CaptureErrorAndWait(err, nil)
		rr := r.Err.(map[string]interface{})
		return errors.New(fmt.Sprintf("(%v) %s", rr["code"].(float64), rr["message"].(string)))
	}

	return nil
}
