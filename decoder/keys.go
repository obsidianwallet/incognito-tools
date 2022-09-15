package decoder

import (
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

func DecodeWalletKey(key string) (interface{}, error) {
	var result interface{}
	wl, err := wallet.Base58CheckDeserialize(v)
	if err != nil {
		panic(err)
	}
	return result, nil
}

// github.com/gin-contrib/gzip v0.0.5
// github.com/gin-gonic/gin v1.8.1
// github.com/go-resty/resty/v2 v2.7.0
// github.com/incognitochain/go-incognito-sdk-v2 v1.0.1-beta.0.20220803110223-48128c589460
// github.com/patrickmn/go-cache v2.1.0+incompatible
