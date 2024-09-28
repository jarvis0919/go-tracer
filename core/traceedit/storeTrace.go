package traceedit

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func StoreTrace(txHash common.Hash, trace interface{}) error {
	jsonData, err := json.MarshalIndent(trace, "", "  ")
	if err != nil {
		log.Fatalf("JSON 编码失败: %v", err)
		return err
	}
	file, err := os.Create("./cache/" + txHash.Hex() + "trace_result.json")
	if err != nil {
		log.Fatalf("无法创建文件: %v", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("无法写入文件: %v", err)
		return err
	}

	return nil
}
