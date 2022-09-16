package decoder

func GetShardIDFromLastByte(b byte, shardNum int) byte {
	return byte(int(b) % shardNum)
}
