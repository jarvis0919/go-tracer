package model

type StructLogs struct {
	Pc      int      `json:"pc"                yaml:"pc"`
	Op      string   `json:"op"                yaml:"op"`
	Gas     int      `json:"gas"               yaml:"gas"`
	Depth   int      `json:"depth"             yaml:"depth"`
	GasCost int      `json:"gasCost,omitempty" yaml:"gasCost"`
	Stack   []string `json:"stack"             yaml:"stack"`
	Memory  []string `json:"memory,omitempty"  yaml:"memory"`
	Error   string   `json:"error,omitempty"   yaml:"error"`
}

type ResultTrace struct {
	Failed      bool         `json:"failed"      yaml:"failed"`
	Gas         int          `json:"gas"         yaml:"gas"`
	ReturnValue string       `json:"returnValue" yaml:"returnValue"`
	StructLogs  []StructLogs `json:"structLogs"  yaml:"structLogs"`
}
