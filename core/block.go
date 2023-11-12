package core

import (
	"github.com/gokutheengineer/utxo-chain/pb"
	"github.com/gokutheengineer/utxo-chain/util"
	protobuf "google.golang.org/protobuf/proto"
)

// HashBlockHeader hashes the block header with BLAKE2b
func HashBlockHeader(header *pb.BlockHeader) (headerHash []byte) {
	// serialize the block header
	h, err := protobuf.Marshal(header)
	if err != nil {
		panic(err)
	}
	headerHash = util.HashBlake2b(h)
	return
}
