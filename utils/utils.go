package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Md5StringByNowTime() string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(time.Now().String()))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
