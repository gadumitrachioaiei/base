package base

import "errors"

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

type Decoder struct {
	decodeMap [256]byte
}

func NewDecoder() *Decoder {
	d := Decoder{}
	for i := 0; i < len(d.decodeMap); i++ {
		d.decodeMap[i] = 0xFF
	}
	for i := 0; i < len(encodeStd); i++ {
		d.decodeMap[encodeStd[i]] = byte(i)
	}
	return &d
}

func (d *Decoder) Decode(source string) ([]byte, error) {
	var src []byte
	for _, c := range source {
		src = append(src, d.decodeMap[c])
	}
	// we decode src into dst
	// we start off from first byte and continue
	// we have an error when we use less than 5 bits of last byte and they are not zero
	var result []byte
	dst := make([]byte, 5)
	length := 0
	for i := 0; i < len(src); i++ {
		// we need last 5 bits of each byte, in different positions
		switch i {
		case 0:
			// we move last 5 bits on the first positions
			dst[0] = src[0] << 3
			length = 1
		case 1:
			// we concatenate 3 bits
			dst[0] |= src[1] >> 2
			// we move last 2 bits on first positions
			dst[1] = src[1] << 6
			length = 2
			if i == len(src)-1 && dst[1] > 0 {
				return nil, errors.New("last byte is not correct")
			}
		case 2:
			// we concatenate 5 bits
			dst[1] |= src[2] << 1
		case 3:
			// we concatenate 1 bit
			dst[1] |= src[3] >> 4
			// we move last 4 bits on the first positions
			dst[2] = src[3] << 4
			length = 3
			if i == len(src)-1 && dst[2] > 0 {
				return nil, errors.New("last byte is not correct")
			}
		case 4:
			// we concatenate 4 bits
			dst[2] |= src[4] >> 1
			dst[3] = src[4] << 7
			length = 4
			if i == len(src)-1 && dst[3] > 0 {
				return nil, errors.New("last byte is not correct")
			}
		case 5:
			// we concatenate 5 bits
			dst[3] |= src[5] << 2
		case 6:
			// we concatenate 2 bits
			dst[3] |= src[6] >> 3
			dst[4] = src[6] << 5
			length = 5
			if i == len(src)-1 && dst[4] > 0 {
				return nil, errors.New("last byte is not correct")
			}
		case 7:
			// we concatenate 5 bits
			dst[4] |= src[7]
			result = append(result, dst[:length]...)
			src = src[8:]
			i = -1
			dst = make([]byte, 5)
			length = 0
		}
	}
	result = append(result, dst[:length]...)
	return result, nil
}
