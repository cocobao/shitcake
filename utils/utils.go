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

func TimeSecToString(sec int64) string {
	return timestampToString(sec, "2006-01-02 15:04:05")
}

func timestampToString(sec int64, formate string) string {
	return time.Unix(sec, 0).Format(string(formate))
}
