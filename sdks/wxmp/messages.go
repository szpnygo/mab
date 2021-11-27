package wxmp

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/szpnygo/mab/sdks/wxmp/messages"
)

func (h *WxMpHelper) InitMessage(echostr, signature, timestamp, nonce string) (string, bool) {
	if len(echostr) == 0 {
		return "", false
	}
	s := makeSignature(h.CallBackToken, timestamp, nonce)
	if s == signature {
		return echostr, true
	}

	return "false", false
}

func (h *WxMpHelper) MessageCallback(data []byte) (*messages.WxMpMessageBody, error) {
	var message messages.WxMpMessageBody
	err := xml.Unmarshal(data, &message)

	return &message, err
}

func makeSignature(token, timestamp, nonce string) string {
	list := []string{token, timestamp, nonce}
	sort.Strings(list)
	s := sha1.New()
	_, _ = io.WriteString(s, strings.Join(list, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
