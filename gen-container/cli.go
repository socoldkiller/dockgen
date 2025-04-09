package gen_container

import (
	"github.com/docker/docker/client"
	"log"
	"sync"
)

var (
	globalClient *client.Client
	once         sync.Once
)

func Client() *client.Client {
	once.Do(func() {
		var err error
		if globalClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()); err != nil {
			log.Fatalf("connnect docker API error ,err: %s", err.Error())
		}
	})
	return globalClient
}
