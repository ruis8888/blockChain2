package block

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string

var nodeVersion = 0

type Block struct {
	header      				BlockHeader
	txs							string
	txCounter					int
	hashCurr					Hash
}

type BlockHeader struct {
	version 					int
	hashPrevBlock			    Hash
	hashMerkleRoot				Hash
	time 						time.Time
	bits 						int
	nonce						int
}

//设置当前区块哈希
func (b *Block) SetHashCurr() *Block {
	//生成头部信息的拼接字符串
	headerStr := b.header.Stringify()
	//计算哈希值
	hashHeaderStr := sha256.Sum256([]byte(headerStr))
	b.hashCurr = fmt.Sprintf("%x",hashHeaderStr)
	return b
}

//区块的构造函数
func NewBlock(prevHash Hash,txs string) *Block {
	//实例化block
	b:=&Block{
		header:   BlockHeader{
			version:        nodeVersion,
			hashPrevBlock:  prevHash,
			//hashMerkleRoot: "",
			time:           time.Time{},
			//bits:           0,
			//nonce:          0,
		},
		txs:       txs,
		txCounter: 1,
		//hashCurr:  "",
	}
	//设置当前区块哈希
	b.SetHashCurr()
	return b
}


//头信息的字符串化
func (bh *BlockHeader) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(),    //得到时间戳，nano级
		bh.bits,
		bh.nonce,
		)
}