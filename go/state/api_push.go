package state

func (self *LuaState) PushNil() {
	self.stack.push(nil)
}

func (self *LuaState) PushBoolean(b bool) {
	self.stack.push(b)
}

func (self *LuaState) PushInteger(n int64) {
	self.stack.push(n)
}

func (self *LuaState) PushNumber(n float64) {
	self.stack.push(n)
}

func (self *LuaState) PushString(s string) {
	self.stack.push(s)
}
