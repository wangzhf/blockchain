package blc

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 创建Block结构体
type Block struct {
	// 高度
	Height int64
	// 上一个区块的哈希值
	PrevBlockHash []byte
	// 交易数据
	Data []byte
	// 时间戳
	TimeStamp int64
	// 哈希值：32个字节，64个16进制数
	Hash []byte
	// 难度
	Nonce int64
}

func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	// 创建区块
	block := &Block{height,
		prevBlockHash,
		[]byte(data),
		time.Now().Unix(), nil, 0}
	// 设置哈希值
	// block.SetHash()

	// 调用工作量证明方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func (block *Block) SetHash() {
	// 将高度转为字节数组
	heightBytes := IntToHex(block.Height)
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	timeBytes := []byte(timeString)
	// 拼接所有的属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes,
	}, []byte{})
	// 生成哈希值
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}
