package base65536

const noByte = -1

// Marshal returns the base65536 encoded version of data.
func Marshal(data []byte) string {
	var res string
	var b1 int16
	var b2 int16
	for i := 0; i < len(data); i += 2 {
		b1 = int16(data[i])
		if i+1 < len(data) {
			b2 = int16(data[i+1])
		} else {
			b2 = noByte
		}
		res += string(getBlockStart[b2] + int(b1))
	}
	return res
}
