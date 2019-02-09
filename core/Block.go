package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64 //区块编号
	TimeStamp int64 //区块时间戳
	PreBlockHash string //上个区块哈希值
	Hash string //当前区块哈希值
	Data string //区块数据
}

func calculateHash(b Block) string  {
	blockData := string(b.Index)+string(b.TimeStamp)+b.PreBlockHash+b.Data
	hashInByte := sha256.Sum256([] byte(blockData))  //调用go语言自带的
	return hex.EncodeToString(hashInByte[:])
}

func GenerateNewBlock(preBlock Block,data string) Block {
	newBlock:=Block{}
	newBlock.Index=preBlock.Index+1
	newBlock.PreBlockHash=preBlock.Hash
	newBlock.TimeStamp=time.Now().Unix()
	newBlock.Data=data
	newBlock.Hash=calculateHash(newBlock)

	return newBlock
}

func GenerateGenesisBlock() Block{ //生成创始区块
	preBlock:= Block{}
	preBlock.Index=-1
	preBlock.Hash=""
	return GenerateNewBlock(preBlock,"Genesis Block")
}