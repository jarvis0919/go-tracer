package client

import (
	"context"
	"demo/model"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func (c *TraceClient) GetTrace(txHash common.Hash) (*model.ResultTrace, error) {
	var result model.ResultTrace
	err := c.Client.Client().CallContext(context.Background(), &result, "debug_traceTransaction", txHash, c.TraceCofing)
	if err != nil {
		log.Fatalf("Failed to trace transaction: %v", err)
		return nil, err
	}
	return &result, nil
}
