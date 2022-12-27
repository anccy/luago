package binchunk

type binChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

// Peototype 函数原型 其中Source, LineInfo, Locvars, UpvalueNames 都是调试信息，非必需
type Prototype struct {
	Source          string        // 二进制chunk的来源，文件/标准输入/字符串编译等
	LineDefined     uint32        // 函数开始行
	LastLineDefined uint32        // 函数结束行
	NumParams       byte          // 固定参数的个数
	IsVararg        byte          // 是否有变长参数
	MaxStackSize    byte          // 寄存器数量 栈可以被用作虚拟寄存器
	Code            []uint32      // 指令表，每条指令占4字节
	Constants       []interface{} // 常量, tag + 常量值
	Upvalues        []Upvalue
	Protos          []*Prototype // 子函数原型列表
	LineInfo        []uint32     // 行号表，和指令表一一对应
	Locvars         []LocVar     // 局部变量，变量名+起止指令索引
	UpvalueNames    []string     // Upvalue名列表，和Upvalues一一对应
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}
