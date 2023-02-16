package state

import (
	"fmt"
	. "github.com/anccy/luago/go/api"
)

func (self *LuaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

func (self *LuaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (self *LuaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

func (self *LuaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

func (self *LuaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

func (self *LuaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

func (self *LuaState) IsTable(idx int) bool {
	return self.Type(idx) == LUA_TTABLE
}

func (self *LuaState) IsFunction(idx int) bool {
	return self.Type(idx) == LUA_TFUNCTION
}

func (self *LuaState) IsThread(idx int) bool {
	return self.Type(idx) == LUA_TTHREAD
}

func (self *LuaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

func (self *LuaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *LuaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *LuaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

func (self *LuaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *LuaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *LuaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *LuaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

func (self *LuaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *LuaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)

	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
