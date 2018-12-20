package main

import (
	"blc"
	"fmt"
)

func main() {
	// 测试Block
	//block := blc.NewBlock("I am a block", make([]byte, 32, 32), 1)
	//fmt.Printf("Height: %x\n", block.Height)
	//fmt.Printf("Data: %s\n", block.Data)

	// 测试创世区块
	//genesisBlock := blc.CreateGenesisBlock("Genesis Block")
	//fmt.Printf("Heigth:%x\n",genesisBlock.Height)
	//fmt.Printf("PrevBlockHash:%x\n",genesisBlock.PrevBlockHash)
	//fmt.Printf("Data:%s\n",genesisBlock.Data)

	// 测试区块链
	//genesisBlockChain := blc.CreateBlockChainWithGenesisBlock("block chain")
	//fmt.Println(genesisBlockChain)
	//fmt.Println(genesisBlockChain.Blocks)
	//fmt.Println(genesisBlockChain.Blocks[0])
	//fmt.Printf("Heigth:%x\n",genesisBlockChain.Blocks[0].Height)
	//fmt.Printf("PrevBlockHash:%x\n",genesisBlockChain.Blocks[0].PrevBlockHash)
	//fmt.Printf("Data:%s\n",genesisBlockChain.Blocks[0].Data)

	// 测试添加新区块
	//blockChain := blc.CreateBlockChainWithGenesisBlock("Genesis Block...")
	//blockChain.AddBlockToBlockChain("Send 100RMB to zhangsan",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 100RMB to lisi",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlockToBlockChain("Send 100RMB to wangwu",
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//for _, block := range blockChain.Blocks {
	//	fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	pow := blc.NewProofOfWork(block)
	//	fmt.Printf("Pow: %v\n", pow.IsValid())
	//	fmt.Println()
	//}

	// 检测pow
	// 创建一个big对象 0000 0000 0000 ... 0001
	//target := big.NewInt(1)
	//fmt.Printf("0x%x\n", target)
	//
	//// 左移256-bit位
	//target = target.Lsh(target, 256 - blc.TargetBit)
	//fmt.Printf("0x%x\n", target)
	//
	//s1 := "hello world"
	//hash := sha256.Sum256([]byte(s1))
	//fmt.Printf("0x%x\n", hash)

	// 测试创世区块存入数据库
	blockchain := blc.CreateBlockChainWithGenesisBlock("Genesis Block...")
	fmt.Println(blockchain)
	defer blockchain.DB.Close()
	// 测试新添加的区块
	blockchain.AddBlockToBlockChain("Send 100 to zhangsan")
	blockchain.AddBlockToBlockChain("Send 200 to lisi")
	blockchain.AddBlockToBlockChain("Send 300 to wangwu")
	fmt.Println(blockchain)
	blockchain.PrintChains()

}
