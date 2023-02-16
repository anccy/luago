package state

type luaStack struct {
	slots []luaValue
	top   int
}

func NewLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

func (ls *luaStack) check(n int) {
	space := len(ls.slots) - ls.top
	for i := space; i < n; i++ {
		ls.slots = append(ls.slots, nil)
	}
}

func (ls *luaStack) push(val luaValue) {
	if ls.top >= len(ls.slots) {
		panic("stack overflow")
	}

	ls.slots[ls.top] = val
	ls.top++
}

func (ls *luaStack) pop() luaValue {
	if ls.top <= 0 {
		panic("stack lowflow")
	}
	ls.top--
	v := ls.slots[ls.top]
	ls.slots[ls.top] = nil
	return v
}

func (ls *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return ls.top + idx + 1
}

func (ls *luaStack) isValid(idx int) bool {
	i := ls.absIndex(idx)
	return i >= 1 && i <= ls.top
}

func (ls *luaStack) get(idx int) luaValue {
	i := ls.absIndex(idx)
	if i <= 0 || i > ls.top {
		return nil
	}
	return ls.slots[idx-1]
}

func (ls *luaStack) set(idx int, val luaValue) {
	i := ls.absIndex(idx)
	if i <= 0 || i > ls.top {
		panic("invalid index")
	}
	ls.slots[i-1] = val
}
