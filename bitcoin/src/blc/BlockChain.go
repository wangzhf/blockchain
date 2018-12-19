package blc

type BlockChain struct {
	// 存储有序的区块
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	// 创世区块
	genesisBlock := CreateGenesisBlock(data)
	// 返回区块链对象
	return &BlockChain{[]*Block{genesisBlock}}
}

func (bc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	// 创建新区块
	newBlock := NewBlock(data, prevHash, height)
	// 添加到切片中
	bc.Blocks = append(bc.Blocks, newBlock)
}
