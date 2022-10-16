package core

import (
	"context"

	"github.com/bee-messenger/core"
	logging "github.com/ipfs/go-log/v2"
	host "github.com/libp2p/go-libp2p/p2p/host/routed"
)

type Bee struct {
	host      host.RoutedHost
	messenger core.Messenger
}

func NewBee(repoPath string) (*Bee, error) {
	err := logging.SetLogLevelRegex(".*", "DEBUG")
	if err != nil {
		panic("logger failed")
	}
	node, err := core.Create(context.Background(), repoPath)
	if err != nil {
		return nil, err
	}
	msgr := core.MessengerBuilder(repoPath+"/msgr_data", node)
	f := &Bee{host: *node, messenger: msgr}

	return f, nil
}
