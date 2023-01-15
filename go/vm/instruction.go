package vm

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

type Instruction uint32

// Opcode 6bit
func (self Instruction) Opcode() int {
	return int(self & 0x3f)
}

func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xff)
	b = int(self >> (6 + 8) & 0x1f)
	c = int(self >> (6 + 8 + 9) & 0x1f)
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
