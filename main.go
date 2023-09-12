package main

import (
	"fmt"

	"github.com/ipoluianov/gazer-billing-contract-eth/api"
)

func main() {
	shop := api.NewShop("./", "http://127.0.0.1:8545", "0x5FbDB2315678afecb367f032d93F642f64180aa3")
	shop.Load()
	shop.Update()
	p := shop.IsPremium("7ctd4bacwqo5xjxypmhblgmffr3t3x5croyetd5n7tsfobnn")
	fmt.Print(p)
}
