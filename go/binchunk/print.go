package binchunk

import "fmt"

func PrintHeader(f *Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}
	varargFlag := ""
	if f.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d, %d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}
