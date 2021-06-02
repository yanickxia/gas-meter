package utils

// GenChecksum gen checksum, input contain magic header, not contains checksum
// 帧头 帧长	目的地址	源地址	命令字	信息段
func GenChecksum(input []byte) byte {
	checksum := 0x0
	// skip magic header
	for i := 2; i < len(input); i++ {
		checksum += int(input[i])
	}

	checksum %= 256
	return byte(256 - checksum)
}
