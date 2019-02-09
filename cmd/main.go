package main

import "core"

func main(){
	bc := core.NewBlockchain()
	bc.SendData("send 1 ETH to Kevin")
	bc.SendData("send 1 BTC to Tom")
	bc.Print()
}