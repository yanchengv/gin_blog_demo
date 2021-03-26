package myutils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	md5 := fmt.Sprintf("%x", h.Sum(nil))
	return md5
}
