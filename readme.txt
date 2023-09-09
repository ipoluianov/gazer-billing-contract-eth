solc --optimize --abi ./contracts/PayloadShop.sol -o build
solc --optimize --bin ./contracts/PayloadShop.sol -o build
abigen --abi=./build/PayloadShop.abi --bin=./build/PayloadShop.bin --pkg=api --out=./api/PayloadShop.go


client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

    