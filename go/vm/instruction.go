package vm

import (
	"fmt"
	"strings"
)

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

type Instruction uint32

// Opcode 6bit
func (self Instruction) Opcode() int {
	return int(self & 0x3f)
}

func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xff)
	c = int(self >> (6 + 8) & 0x1ff)
	b = int(self >> (6 + 8 + 9) & 0x1ff)
	return
}

func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xff)
	bx = int(self >> (6 + 8))
	return
}

func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

func (self Instruction) Ax() int {
	return int(self >> 6)
}

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}

func (self *Instruction) String() string {
	ret := &strings.Builder{}
	_, _ = fmt.Fprintf(ret, "%s\t", self.OpName())
	switch self.OpMode() {
	case IABC:
		a, b, c := self.ABC()
		_, _ = fmt.Fprintf(ret, "%d", a)
		if self.BMode() != OpArgN {
			if b > 0xff {
				_, _ = fmt.Fprintf(ret, " %d", -1-b&0xff)
			} else {
				_, _ = fmt.Fprintf(ret, " %d", b)
			}
		}
		if self.CMode() != OpArgN {
			if c > 0xff {
				_, _ = fmt.Fprintf(ret, " %d", -1-c&0xff)
			} else {
				_, _ = fmt.Fprintf(ret, " %d", c)
			}
		}
	case IABx:
		a, bx := self.ABx()
		_, _ = fmt.Fprintf(ret, "%d", a)
		if self.BMode() == OpArgK {
			_, _ = fmt.Fprintf(ret, " %d", -1-bx)
		} else if self.BMode() == OpArgU {
			_, _ = fmt.Fprintf(ret, " %d", bx)
		}
	case IAsBx:
		a, sbx := self.AsBx()
		_, _ = fmt.Fprintf(ret, "%d %d", a, sbx)
	case IAx:
		ax := self.Ax()
		_, _ = fmt.Fprintf(ret, "%d", -1-ax)
	}
	return ret.String()
}
