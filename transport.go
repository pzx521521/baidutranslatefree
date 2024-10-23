package baidutranslatefree

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type BaiduTranslater struct {
	Signer *Signer
	ID     string
	Token  string
	From   string
	To     string
}

func NewBaiduTranslater() (*BaiduTranslater, error) {
	signer, err := NewSigner()
	if err != nil {
		return nil, err
	}
	baiduid, token, err := GetIDAndToken()
	if err != nil {
		return nil, err
	}
	return &BaiduTranslater{Signer: signer, ID: baiduid, Token: token, From: "en", To: "zh"}, nil
}
func (t *BaiduTranslater) SetFromTo(from, to string) error {
	if _, ok := LangList[from]; !ok {
		return errors.New("from language not support")
	}
	if _, ok := LangList[to]; !ok {
		return errors.New("to language not support")
	}
	t.From = from
	t.To = to
	return nil
}

// TransPort 翻译
func (t *BaiduTranslater) Translate(input string) (text string, err error) {
	client := &http.Client{}
	signature, err := t.Signer.CalSign(input)
	if err != nil {
		return text, err
	}
	param := fmt.Sprintf("query=%s&from=%s&to=%s&token=%s&sign=%s",
		url.PathEscape(input),
		t.From,
		t.To,
		t.Token,
		signature,
	) // 创建请求
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(param))

	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 将 Cookie 拆分成多行以增加可读性
	cookies := []string{
		"BAIDUID=" + t.ID,
	}
	// 连接 Cookie 字符串
	req.Header.Set("Cookie", strings.Join(cookies, "; ")) // 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return text, err
	}
	defer resp.Body.Close()
	var output map[string]interface{}
	// 读取响应
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return text, err
	}
	// 提取 trans 的 dst
	trans, ok := output["trans"].([]interface{})
	if !ok {
		fmt.Println("not found output resp trans")
		return
	}

	// 遍历 trans，提取 dst
	for _, item := range trans {
		if transMap, ok := item.(map[string]interface{}); ok {
			if dst, ok := transMap["dst"].(string); ok {
				return dst, nil
			}
		}
		//不清楚他为什么返回一个数组
	}
	return text, errors.New("not found dst in trans")
}
