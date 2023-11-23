package gaudius

import (
	"context"
	"crypto/ecdsa"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (sdk *AudiusSdk) write(encodedAbi string) (*common.Hash, error) {
	ctx := context.TODO()
	// https://goethereumbook.org/wallet-generate/
    privateKey, err := crypto.GenerateKey()
    if err != nil {
       return nil, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	sender := common.HexToAddress(address)

	// gather tx params
	client := sdk.AcdcClient
	nonce, err := client.NonceAt(ctx, sender, nil)
	if err != nil {
		return nil, err
	}

	em := sdk.EntityManagerAddress

	tx := types.NewTransaction(nonce, *em, nil, 2000000, nil, []byte(encodedAbi))
	hash := tx.Hash()

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	
	return &hash, nil
}

func (sdk *AudiusSdk) RawWrite(encodedAbi string) (*contracts.EntityManagerManageEntity, error) {
	ctx := context.TODO()
	block, err := sdk.AcdcClient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	
	blocknum := block.NumberU64()
	tx, err := sdk.write(encodedAbi)
	if err != nil {
		return nil, err
	}

	sp := &ScannerParams{ StartBlock: &blocknum }
	sub := sdk.TxSubscriber(sp, *tx)

	event := <-sub
	return event, nil
}
