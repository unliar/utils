package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"net/url"
	"sort"
	"strings"
	"fmt"
)

type M map[string]string

// 支付宝签名
func (m M) CommonPublicKeySign(AliPayPublicKey *rsa.PublicKey, AppPrivateKey *rsa.PrivateKey, SignType string) (string, error) {
	var data []string
	for k, v := range m {
		if k != "sign" && v != "" {
			data = append(data, "%s=%s", k, v)
		}
	}
	sort.Strings(data)
	// 待签名字符串
	signStr := strings.Join(data, "&")
	fmt.Println("debug签名字符串=>",signStr)
	// 默认是RSA2 256
	s, cs := GetSignOpsBySignType(SignType)
	_, err := s.Write([]byte(signStr))
	if err != nil {
		return "", err
	}
	hashByte := s.Sum(nil)
	SignByte, err := AppPrivateKey.Sign(rand.Reader, hashByte, cs)
	if err != nil {
		return "", err
	}
	return url.QueryEscape(base64.StdEncoding.EncodeToString(SignByte)), nil
}

// 根据签名的类型获取不同的值
func GetSignOpsBySignType(t string) (hash.Hash, crypto.SignerOpts) {
	if t == "RSA1" {
		return sha1.New(), crypto.SHA1
	}
	return sha256.New(), crypto.SHA256
}
