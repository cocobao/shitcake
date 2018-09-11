package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
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

func StructToMapJson(src interface{}) map[string]interface{} {
	b, _ := json.Marshal(src)

	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}

func isFileExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
