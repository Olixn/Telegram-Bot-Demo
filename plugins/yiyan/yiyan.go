package yiyan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetYiYan() (str string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://v1.hitokoto.cn", strings.NewReader(""))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.62")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var respond map[string]interface{}
	err = json.Unmarshal(body, &respond)
	if err != nil {
		log.Println(err)
		return
	}
	return respond["hitokoto"].(string), nil
}
