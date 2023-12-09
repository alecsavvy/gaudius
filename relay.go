package gaudius

import (
	"context"
	"crypto/ecdsa"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (sdk *AudiusSdk) relay(encodedabi string) (*contracts.EntityManagerManageEntity, error) {
	txhash, err := sdk.genTxHash(encodedabi)
	if err != nil {
		return nil, err
	}
	sub := sdk.TxSubscriber(nil, *txhash)
	return <-sub, nil
}

func (sdk *AudiusSdk) genTxHash(encodedAbi string) (*common.Hash, error) {
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
