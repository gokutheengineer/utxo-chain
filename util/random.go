package util

import (
	"crypto/rand"

	crypto "github.com/gokutheengineer/utxo-chain/crypto"
	"github.com/gokutheengineer/utxo-chain/pb"
)

func randomTransactionInput() *pb.TransactionInput {
	inputSK, err := crypto.NewPrivateKey()
	if err != nil {
		panic(err)
	}
	inputPK := inputSK.GeneratePublicKey()

	// signature will be computed over the transaction data, after the transaction is assembled.
	// for each input being signed, modify tx data a bit.
	// for example, for the current input set its output to the script pub key of the output spending it (for P2PKH)
	// serialize the transaction with these modifications
	// transaction data will be hashed with blake2b to create a digest.
	// the digest will be signed with the private key corresponding to the public key in ScriptSig
	// replace the signature int the ScriptSig.
	// repeat for each input. the transaction is ready to be broadcasted.
	scriptSig := &pb.ScriptSig{
		Type:         pb.ScriptType_P2PKH,
		Signature:    nil,
		PublicKey:    inputPK.Bytes(),
		CustomScript: nil,
	}

	return &pb.TransactionInput{
		PrevOutTxHash: HashBlake2b(randomBytes(32)),
		PrevOutIndex:  0,
		ScriptSig:     scriptSig,
		Sequence:      HashBlake2b(randomBytes(32)),
	}
}

func randomTransactionOutput() *pb.TransactionOutput {
	return &pb.TransactionOutput{
		Address: []byte{'a', 'b', 'c'},
		Amount:  1,
	}
}

func RandomTranscationSingleInputOutput() *pb.Transaction {
	txInputs := []*pb.TransactionInput{randomTransactionInput()} //make([]*pb.TransactionInput, 0)
	txOutputs := []*pb.TransactionOutput{randomTransactionOutput()}
	data := randomBytes(32)
	dataHash := HashBlake2b(data)

	return &pb.Transaction{
		Inputs:              txInputs,
		Outputs:             txOutputs,
		Fee:                 0,
		Data:                data,
		TransactionDataHash: dataHash,
	}
}

func RandomBlockHeader() *pb.BlockHeader {
	return &pb.BlockHeader{}
}

func RandomBlock() *pb.Block {
	return &pb.Block{}
}

func randomBytes(size int) []byte {
	random := make([]byte, size)
	_, err := rand.Read(random)
	if err != nil {
		panic(err)
	}
	return random
}
