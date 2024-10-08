package main

import (
	"fmt"
	"log"

	"github.com/jarvis0919/go-tracer/core/client"
	"github.com/jarvis0919/go-tracer/core/traceedit"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	c, err := client.Defult()
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	txHash := common.HexToHash("0xed2e82bb59e2ea39bfdc7d08ae2f7fcad7200e00163f6e3440b9a5d72fc3ef5d")
	trace, err := c.GetTrace(txHash)
	if err != nil {
		log.Fatalf("Failed to trace transaction: %v", err)
	}

	//TODO:增加数据处理工作

	traceedit.CookTrace(trace.StructLogs, false)
	traceedit.StoreTrace(txHash, trace.StructLogs)

	fmt.Println(trace.Gas)
	fmt.Println(len(trace.StructLogs))

}
