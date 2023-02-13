package main

import (
	"github.com/anccy/luago/go/binchunk"
)

func main() {
	proto := binchunk.ParseChunkFile("./lua/luac.out")
	binchunk.List(proto)
}
