package binchunk

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

type binChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

// lua: DumpHeader
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

// lua: DumpFunction
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

type reader struct {
	data []byte
}

func (r *reader) readByte() byte {
	b := r.data[0]
	r.data = r.data[1:]
	return b
}

func (r *reader) readUint32() uint32 {
	ret := binary.LittleEndian.Uint32(r.data)
	r.data = r.data[4:]
	return ret
}

func (r *reader) readUint64() uint64 {
	ret := binary.LittleEndian.Uint64(r.data)
	r.data = r.data[8:]
	return ret
}

func (r *reader) readLuaInteger() int64 {
	return int64(r.readUint64())
}

func (r *reader) readLuaNumber() float64 {
	return math.Float64frombits(r.readUint64())
}

func (r *reader) readString() string {
	size := uint(r.readByte())
	if size == 0 {
		return ""
	} else if size == 0xff { // long string
		size = uint(r.readUint64())
	}
	bytes := r.readBytes(size - 1)
	return string(bytes)
}

func (r *reader) readBytes(n uint) []byte {
	ret := r.data[:n]
	r.data = r.data[n:]
	return ret
}

func (r *reader) checkHeader() error {
	if string(r.readBytes(4)) != LUA_SIGNATURE {
		return errors.New("signature")
	}
	if r.readByte() != LUAC_VERSION {
		return errors.New("version")
	}
	if r.readByte() != LUAC_FORMAT {
		return errors.New("format")
	}
	if string(r.readBytes(6)) != LUAC_DATA {
		return errors.New("corrupted")
	}
	if r.readByte() != CINT_SIZE {
		return errors.New("int size")
	}
	if r.readByte() != CSIZET_SIZE {
		return errors.New("sizet size")
	}
	if r.readByte() != INSTRUCTION_SIZE {
		return errors.New("instruction size")
	}
	if r.readByte() != LUA_INTEGER_SIZE {
		return errors.New("lua integer size")
	}
	if r.readByte() != LUA_NUMBER_SIZE {
		return errors.New("lua number size")
	}
	if r.readLuaInteger() != LUAC_INT {
		return errors.New("luac int")
	}
	if r.readLuaNumber() != LUAC_NUM {
		return errors.New("luac num")
	}
	return nil
}

func ParseChunk(s string) {
	rd := &reader{
		data: []byte(s),
	}
	fmt.Println(rd.checkHeader())
	//for {
	//	b := rd.readByte()
	//	fmt.Printf("%c", b)
	//}
}

func ParseChunkFile(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ParseChunk(string(b))
}

func Cmd() {
	ParseChunkFile("./lua/luac.out")
}

// ANCCY TODO
// A Luac parser
