package traceedit

import (
	"demo/global"
	"demo/model"
	"fmt"
)

func CookTrace(logs []model.StructLogs, FullTrace bool) []model.StructLogs {
	lastOpcode := make(map[int]string)

	for k, v := range logs {
		depth := v.Depth
		if _, ok := lastOpcode[depth]; !ok {
			logs[k] = pruneStructLog(v, "", FullTrace)
		} else {
			logs[k] = pruneStructLog(v, lastOpcode[depth], FullTrace)
		}
		lastOpcode[depth] = logs[k].Op
	}
	return logs

}

func pruneStructLog(log model.StructLogs, lastOpcode string, FullTrace bool) model.StructLogs {
	prune1 := true
	prune2 := true
	if FullTrace {
		prune1 = false
		prune2 = false
	}
	structLog_copy := log
	// Prune 1: remove pc, and gasCost
	if prune1 {
		structLog_copy.GasCost = 0
	}

	// Prune 2: remove unnessary stack (won't be used by opcode)
	if prune2 {
		var len1 int
		var len2 int = 0
		var ok bool
		if len1, ok = global.OpcodeInputStackmap[structLog_copy.Op]; !ok {
			panic(fmt.Sprint("Opcode %s not found in opcode2InputStackLength!!\n", structLog_copy.Op))
		}

		if lastOpcode != "" {
			if len2, ok = global.OpcodeOutputStackmap[lastOpcode]; !ok {
				panic(fmt.Sprint("Opcode %s not found in opcode2OutputStackLength!!\n", lastOpcode))
			}
		}
		necessaryStackLen := max(len1, len2)
		if structLog_copy.Pc == 236 {
			fmt.Println(structLog_copy.Stack, len(structLog_copy.Stack)-necessaryStackLen)
		}
		if len(structLog_copy.Stack) != 0 {
			start := len(structLog_copy.Stack) - necessaryStackLen%len(structLog_copy.Stack)
			structLog_copy.Stack = structLog_copy.Stack[start:]
		}

		// for k, v := range structLog_copy.Stack {
		// 	structLog_copy.Stack[k] = fmt.Sprintf("0x%x", v)
		// }
	}

	// Prune 3: remove unnessary memory (won't be used by opcode)
	var matchSuccessful bool
	switch {
	case structLog_copy.Op == "RETURN", structLog_copy.Op == "REVERT", structLog_copy.Op == "KECCAK256", structLog_copy.Op == "CODECOPY", structLog_copy.Op == "EXTCODECOPY", structLog_copy.Op == "RETURNDATACOPY", structLog_copy.Op == "SHA3":
		matchSuccessful = true
	case structLog_copy.Op == "CREATE", structLog_copy.Op == "CREATE2", structLog_copy.Op == "CALL", structLog_copy.Op == "CALLCODE", structLog_copy.Op == "DELEGATECALL", structLog_copy.Op == "STATICCALL":
		matchSuccessful = true
	case structLog_copy.Op == "RETURN", structLog_copy.Op == "REVERT", structLog_copy.Op == "SELFDESTRUCT", structLog_copy.Op == "STOP", structLog_copy.Op == "INVALID":
		matchSuccessful = true
	case lastOpcode == "CALLDATACOPY", lastOpcode == "CODECOPY", lastOpcode == "EXTCODECOPY", lastOpcode == "RETURNDATACOPY":
		lastOpcode = "RETURNDATACOPY"
		matchSuccessful = true
	case lastOpcode == "CALL", lastOpcode == "CALLCODE", lastOpcode == "STATICCALL", lastOpcode == "DELEGATECALL":
		matchSuccessful = true
	}

	if !matchSuccessful && structLog_copy.Memory != nil {
		structLog_copy.Memory = nil
	}

	// res := utils.StructToMap(structLog_copy)
	// if structLog_copy.Error == "" {
	// 	delete(res, "error")
	// }
	// if structLog_copy.GasCost == -1 {
	// 	delete(res, "gasCost")
	// }
	// if structLog_copy.Memory == nil {
	// 	delete(res, "memory")
	// }
	return structLog_copy

}
