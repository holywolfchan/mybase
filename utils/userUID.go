package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/holywolfchan/yuncang/model"
)

var secret = "mannabrain"

func GetUserUID(u *model.User) (s string, err error) {
	if strings.TrimSpace(u.Passport) != "" {
		s := u.Passport + secret + u.Password
		r := Sha1([]byte(s))
		return r, nil
	}
	return "", fmt.Errorf("userpassport should not be empty,current user info:%v", u)
}

func Sha1(b []byte) string {
	r := sha1.Sum(b)
	return hex.EncodeToString(r[:])
}
