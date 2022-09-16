package decoder

import (
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/common/base58"
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
	"github.com/obsidianwallet/incognito-tools/shared"
)

func DecodeWalletKey(key string, shardNum int) (*shared.DecodedWallet, error) {
	wl, err := wallet.Base58CheckDeserialize(key)
	if err != nil {
		return nil, err
	}

	otakey := wl.Base58CheckSerialize(wallet.OTAKeyType)

	readonlykey := wl.Base58CheckSerialize(wallet.ReadonlyKeyType)

	paymentkey := wl.Base58CheckSerialize(wallet.PaymentAddressType)

	validatorkey := ""
	if len(wl.KeySet.PrivateKey) > 0 {
		validatorkey = base58.Base58Check{}.Encode(common.HashB(common.HashB(wl.KeySet.PrivateKey)), common.ZeroByte)
	}

	shardid := -1
	var pubkey string
	if len(wl.KeySet.PaymentAddress.Pk) > 0 {
		pubkey, err = wl.GetPublicKey()
		if err != nil {
			return nil, err
		}
		shardid = int(GetShardIDFromLastByte(wl.KeySet.PaymentAddress.Pk[len(wl.KeySet.PaymentAddress.Pk)-1], shardNum))
	}

	result := shared.DecodedWallet{
		PubkeyBase58:   pubkey,
		PaymentAddress: paymentkey,
		OTAKeyBase58:   otakey,
		ReadOnlyKey:    readonlykey,
		ValidatorKey:   validatorkey,
		ShardID:        shardid,
	}

	return &result, nil
}
