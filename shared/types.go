package shared

type DecodedWallet struct {
	PubkeyBase58   string
	PaymentAddress string
	OTAKeyBase58   string
	ReadOnlyKey    string
	ValidatorKey   string
	ShardID        int
}
