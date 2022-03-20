package innerLink

import (
	//"github.com/i-Things/things/src/ddsvr/internal/domain"
	"context"
	"github.com/i-Things/things/src/ddsvr/ddDef"
	"github.com/i-Things/things/src/ddsvr/internal/config"
)

type (
	InnerLink interface {
		Publish(topic string, payload []byte) error
		Subscribe(handle Handle) error
	}
	Handle         func(ctx context.Context) InnerSubHandle
	InnerSubHandle interface {
		Publish(info *ddDef.InnerPublish) error
		//Login(ctx context.Context, out *domain.LogInOut) error
		//Logout(ctx context.Context, out *domain.LogInOut) error
	}
)

func NewInnerLink(conf config.InnerLinkConf) (InnerLink, error) {
	return NewNatsClient(conf.Nats)
}
