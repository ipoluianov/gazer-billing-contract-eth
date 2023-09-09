package api

import (
	"encoding/base32"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Shop struct {
	lastId  int64
	records []ShopRecord

	dbPath          string
	connectionPoint string
	contractAddress string
}

type ShopRecord struct {
	id      int64
	address string
	payload string
}

func NewShop(dbPath string, connectionPoint string, contractAddress string) *Shop {
	var c Shop
	c.dbPath = dbPath
	c.connectionPoint = connectionPoint
	c.contractAddress = contractAddress
	return &c
}

func (c *Shop) parseAndAddLine(line string) {
	line = strings.Trim(line, " \r\n")
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == ';'
	})
	if len(parts) != 3 {
		return
	}

	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return
	}

	ethAddress := parts[1]
	xchgAddress := parts[2]

	var rec ShopRecord
	rec.id = id
	rec.address = ethAddress
	rec.payload = xchgAddress
	c.records = append(c.records, rec)

	if rec.id > c.lastId {
		c.lastId = rec.id
	}
	fmt.Println("loaded row", line)
}

func (c *Shop) addRecord(id int64, ethAddress string, xchgAddressBS []byte) error {
	if len(xchgAddressBS) != 32 {
		return errors.New("wrong address format")
	}
	if id <= c.lastId {
		return errors.New("already exists")
	}
	xchgAddressBS = xchgAddressBS[:30]

	xchgAddress := strings.ToLower(base32.StdEncoding.EncodeToString(xchgAddressBS))

	var r ShopRecord
	r.id = id
	r.address = ethAddress
	r.payload = xchgAddress
	c.records = append(c.records, r)
	err := c.appendLineToFile(r)
	if err != nil {
		return err
	}
	if r.id > c.lastId {
		c.lastId = r.id
	}
	return nil
}

func (c *Shop) appendLineToFile(rec ShopRecord) error {
	file, err := os.OpenFile(c.fileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = file.WriteString(strconv.FormatInt(rec.id, 10) + ";" + rec.address + ";" + rec.payload + "\r\n")
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	return nil
}

func (c *Shop) fileName() string {
	filename := c.dbPath + "/" + c.contractAddress + ".txt"
	return filename
}

func (c *Shop) Load() error {
	bs, err := os.ReadFile(c.fileName())
	if err != nil {
		return err
	}

	currentLine := make([]byte, 0)
	for i := 0; i < len(bs); i++ {
		if bs[i] == 13 || bs[i] == 10 {
			c.parseAndAddLine(string(currentLine))
			currentLine = make([]byte, 0)
			continue
		}
		currentLine = append(currentLine, bs[i])
	}
	c.parseAndAddLine(string(currentLine))
	return nil
}

func (c *Shop) Update() error {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return err
	}

	contract, err := NewApiCaller(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), client)
	if err != nil {
		return err
	}

	recordsCount, err := contract.RecordsCount(nil)
	if err != nil {
		return err
	}

	fmt.Println("recordsCount", recordsCount)
	for i := int64(c.lastId + 1); i < recordsCount.Int64()+1; i++ {
		fmt.Println("Processing id =", i)
		rec, err := contract.Records(nil, big.NewInt(i))
		if err != nil {
			return err
		}
		err = c.addRecord(i, rec.Owner.Hex(), rec.Payload[:])
		if err != nil {
			fmt.Println("ADD RECORD ERROR:", err)
		}
	}

	client.Close()
	return nil
}

func (c *Shop) GetRecords() {
}
