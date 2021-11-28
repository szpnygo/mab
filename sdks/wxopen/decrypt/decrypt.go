package decrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"

	"github.com/szpnygo/mab/internal/errorx"
)

type WXBizMsgCrypt struct {
	Token          string
	EncodingAesKey string
	AppId          string
}

func (wmc *WXBizMsgCrypt) Decrypt(encrypted string) ([]byte, error) {
	var data []byte
	var err error
	var key []byte
	key, err = base64.StdEncoding.DecodeString(wmc.EncodingAesKey + "=")
	if err != nil {
		return nil, errorx.Error(err)
	}

	data, err = base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, errorx.Error(err)
	}

	var c cipher.Block
	c, err = aes.NewCipher(key)
	if err != nil {
		return nil, errorx.Error(err)
	}

	cbc := cipher.NewCBCDecrypter(c, []byte(wmc.EncodingAesKey[:16]))
	cbc.CryptBlocks(data, data)

	decoded := PKCS7Decode(data)

	buf := bytes.NewBuffer(decoded[16:20])

	var msgLen int32
	_ = binary.Read(buf, binary.BigEndian, &msgLen)

	msgDecrypt := decoded[20 : 20+msgLen]

	return msgDecrypt, errorx.Error(err)
}

func PKCS7Decode(text []byte) []byte {
	pad := int(text[len(text)-1])

	if pad < 1 || pad > 32 {
		pad = 0
	}

	return text[:len(text)-pad]
}

func PKCS7Encode(text []byte) []byte {
	const BlockSize = 32

	amountToPad := BlockSize - len(text)%BlockSize

	for i := 0; i < amountToPad; i++ {
		text = append(text, byte(amountToPad))
	}

	return text
}
