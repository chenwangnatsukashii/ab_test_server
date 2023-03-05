package line_http

import (
	"encoding/json"
	"fmt"
	"io"
	"line_china/search_server/src/model"
	"net/http"
	"strings"
	"time"
)

// HttpGet get方法
func HttpGet(url string) (map[string]interface{}, int) {
	var list model.HttpData

	list.Host = "GET"
	list.Url = "http://127.0.0.1:8081/line_china/abtest" + url
	req, err := http.NewRequest(list.Host, list.Url, strings.NewReader(list.Data))
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	ContentType := "application/json"
	Accept := "*/*"
	AcceptEncoding := "gzip, deflate, br"
	UserAgent := list.Data
	Connection := "keep-alive"

	req.Header.Add("Content-Length", list.Url)
	req.Header.Add("Host", list.Host)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", Accept)
	req.Header.Add("Accept-Encoding", AcceptEncoding)
	req.Header.Add("Connection", Connection)
	req.Header.Add("Content-Type", ContentType)

	client := &http.Client{}
	var startTime, costTime int
	startTime = int(time.Now().UnixNano() / 1e6)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	lin, _ := io.ReadAll(resp.Body)
	endTime := int(time.Now().UnixNano() / 1e6)
	costTime = endTime - startTime
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Expect 200, %d GET!", resp.StatusCode)
	}

	classDetailMap := make(map[string]interface{})
	err = json.Unmarshal(lin, &classDetailMap)

	return classDetailMap, costTime

}
