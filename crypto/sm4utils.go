package crypto

import (
	"errors"
	"encoding/base64"
	"sync"
)

type SM4Utils struct {
}

var s *SM4Utils
var once sync.Once

func NewSM4Utils() *SM4Utils {
	once.Do(func() {
		s = &SM4Utils{}
	})
	return s
}

func (this *SM4Utils) GetEncStr(input, key *string) (string, error) {
	if s, err := this.sm4encDec([]byte(*input), []byte(*key), true); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(s), nil
	}
}

func (this *SM4Utils) GetDecStr(input, key *string) (string, error) {
	inputBytes, err := base64.StdEncoding.DecodeString(*input)
	if err != nil {
		return "", err
	}
	if s, err := this.sm4encDec(inputBytes, []byte(*key), false); err != nil {
		return "", err
	} else {
		return string(s), nil
	}
}

func (this *SM4Utils) sm4encDec(input, key []byte, isEnc bool) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("key length err")
	}
	keyBytes := key
	var dataBytes []byte
	if isEnc {
		dataBytes = this.encPading(input)
	} else {
		dataBytes = input
	}
	if len(dataBytes)%16 != 0 {
		return nil, errors.New("Pading err:length err")
	}
	c, err := NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	var outstr []byte
	for i := 0; i < len(dataBytes); i += 16 {
		out := make([]byte, 16)
		in := dataBytes[i:i+16]
		if isEnc {
			c.Encrypt(out, in)
		} else {
			c.Decrypt(out, in)
		}
		outstr = append(outstr, out...)
	}
	if !isEnc {
		outstr = this.decPading(outstr)
	}
	return outstr, nil
}

func (this *SM4Utils) encPading(data []byte) []byte {
	p := 16 - len(data)%16
	dataBytes := data
	for i := 0; i < p; i++ {
		dataBytes = append(dataBytes, byte(p))
	}
	return dataBytes
}

func (this *SM4Utils) decPading(data []byte) []byte {
	p := data[len(data)-1]
	return data[:(len(data) - int(p))]
}
