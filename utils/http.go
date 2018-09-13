package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cocobao/log"
	"golang.org/x/net/context/ctxhttp"
)

func DoHttpPostSimple(url string, bd interface{}) ([]byte, error) {
	var req *http.Request
	var res *http.Response
	var err error

	ctx := context.Background()
	client := &http.Client{}
	tryTime := 0

	var data []byte
	if bd != nil {
		data, _ = json.Marshal(bd)
	} else {
		data, _ = json.Marshal(map[string]interface{}{})
	}

tryAgain:
	req, err = http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	ctxto, cancel := context.WithTimeout(ctx, 3*time.Second)
	res, err = ctxhttp.Do(ctxto, client, req)
	cancel()
	if err != nil {
		log.Warn("push post err:", err, tryTime)
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}
		//最多重试3次
		tryTime++
		if tryTime < 3 {
			goto tryAgain
		}
		return nil, err
	}
	if res.Body == nil {
		return nil, errors.New("post response is nil")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("post result:%d", res.StatusCode)
	}
	var result []byte
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
