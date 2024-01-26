package usecase

import (
	"github.com/centrifugal/gocent/v3"
	"golang.org/x/net/context"
)

type CentrifugoAPI struct {
	Api *gocent.Client
}

func (c *CentrifugoAPI) Publish(channel string, data []byte) (gocent.PublishResult, error) {

	ctx := context.Background()
	return c.Api.Publish(ctx, channel, data)

}
