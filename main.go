package main

import (
	"github.com/demochain/core"
)

func main() {
	bc := core.NewBlockchain()
	bc.SendData("this is one")
	bc.SendData("this is two")
	bc.SendData("this is three")
	bc.SendData("this is four")
	bc.Print()
}
