package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type TraceClient struct {
	*ethclient.Client
	TraceCofing TraceCofing
}

type TraceCofing struct {
	enableMemory     bool
	disableMemory    bool
	disableStack     bool
	disableStorage   bool
	enableReturnData bool
}

func Defult() (*TraceClient, error) {
	Client, err := ethclient.Dial("https://eth-pokt.nodies.app/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return &TraceClient{
		Client,
		TraceCofing{
			enableMemory:     true,
			disableMemory:    false,
			disableStack:     false,
			disableStorage:   true,
			enableReturnData: false,
		},
	}, nil
}

func NewTraceClient(traceCofing TraceCofing) *TraceClient {
	Client, err := ethclient.Dial("https://eth-pokt.nodies.app/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil
	}
	return &TraceClient{
		Client,
		traceCofing,
	}
}
