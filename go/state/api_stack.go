package state

func (self *LuaState) GetTop() int {
	return self.stack.top
}

func (self *LuaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *LuaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

func (self *LuaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

func (self *LuaState) Copy(from, to int) {
	v := self.stack.get(from)
	self.stack.set(to, v)
}

func (self *LuaState) PushValue(idx int) {
	v := self.stack.get(idx)
	self.stack.push(v)
}

func (self *LuaState) Replace(idx int) {
	v := self.stack.pop()
	self.stack.set(idx, v)
}

func (self *LuaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

func (self *LuaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

func (self *LuaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}

func (self *LuaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow")
	}

	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
