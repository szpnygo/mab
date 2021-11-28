package wxopen

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/sdks/wxopen/decrypt"
	"github.com/szpnygo/mab/sdks/wxopen/messages"
)

type EncryptMsg struct {
	XMLName xml.Name `xml:"xml"`
	AppId   string   `xml:"AppId"`
	Encrypt string   `xml:"Encrypt"`
}

func (h *WxOpenHelper) MessageCallback(data []byte, msgSignature string, timestamp string, nonce string) (*messages.WxOpenMessageBody, error) {
	var message EncryptMsg
	err := xml.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}

	selfSign := getSHA1(h.CallbackToken, timestamp, nonce, string(data))
	if selfSign != msgSignature {
		return nil, errorx.New("signature error")
	}

	wxCrypt := decrypt.WXBizMsgCrypt{
		Token:          h.CallbackToken,
		EncodingAesKey: h.EncodingAESKey,
		AppId:          h.ComponentAppID,
	}
	result, err := wxCrypt.Decrypt(message.Encrypt)
	if err != nil {
		return nil, err
	}

	var event *messages.WxOpenMessageBody
	err = xml.Unmarshal([]byte(result), event)
	if err != nil {
		return nil, err
	}

	return event, nil

}

func getSHA1(token string, timestamp string, nonce string, encrypt_msg string) string {
	list := []string{encrypt_msg, token, timestamp, nonce}
	sort.Strings(list)
	s := sha1.New()
	_, _ = io.WriteString(s, strings.Join(list, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
