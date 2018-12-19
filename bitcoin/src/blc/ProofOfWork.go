package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 0000 0000 0000 0000 1001 0001 0000 ... 0001
// 256位Hash里面前面至少有16个零
const TargetBit = 16

type ProofOfWork struct {
	// 要验证的区块
	Block *Block

	// 大整数存储，目标哈希值
	Target *big.Int
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	// 创建一个big对象： 0000 0000....0001
	target := big.NewInt(1)

	// 左移256-bit位
	target = target.Lsh(target, 256-TargetBit)
	return &ProofOfWork{block, target}
}

// 根据Block生成一个byte数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(TargetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// 返回有效的哈希和nonce值
func (pow *ProofOfWork) Run() ([]byte, int64) {
	// 1. 将Block的属性拼接成字节数组
	// 2. 生成Hash
	// 3. 循环判断Hash的有效性，满足条件，跳出循环结束验证
	nonce := 0
	hashInt := new(big.Int)
	var hash [32]byte
	for {
		// 获取字节数组
		dataBytes := pow.prepareData(nonce)
		// 生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%d: %x", nonce, hash)
		// 将Hash存储到hashInt
		hashInt.SetBytes(hash[:])
		// 判断hashInt是否小于Block里的target
		if pow.Target.Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	return hash[:], int64(nonce)
}

// 验证
func (pow *ProofOfWork) IsValid() bool {
	hashInt := new(big.Int)
	hashInt.SetBytes(pow.Block.Hash)
	return pow.Target.Cmp(hashInt) == 1
}
