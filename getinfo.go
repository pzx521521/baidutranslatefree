package baidutranslatefree

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// GetCookies 发起请求获取 fanyi.baidu.com 的 Cookies
func GetCookies() ([]*http.Cookie, error) {
	// 1. 创建请求
	req, err := http.NewRequest("GET", "https://fanyi.baidu.com", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 2. 设置自定义的 User-Agent 头
	req.Header.Set("User-Agent", userAgent)

	// 3. 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching page: %w", err)
	}
	defer resp.Body.Close()

	// 4. 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// 5. 返回 Cookies
	return resp.Cookies(), nil
}

// GetIDAndToken 发起请求获取 BAIDUID 和 token，设置自定义 User-Agent
func GetIDAndToken() (baiduid string, token string, err error) {
	// 1. 调用 GetCookies 获取 Cookies
	cookies, err := GetCookies()
	if err != nil {
		return "", "", fmt.Errorf("error getting cookies: %w", err)
	}

	// 2. 创建请求
	req, err := http.NewRequest("GET", "https://fanyi.baidu.com", nil)
	if err != nil {
		return "", "", fmt.Errorf("error creating request: %w", err)
	}

	// 3. 设置自定义的 User-Agent 头
	req.Header.Set("User-Agent", userAgent)

	// 4. 将 Cookies 附加到请求中
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// 5. 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error fetching page: %w", err)
	}
	defer resp.Body.Close()

	// 6. 提取 Set-Cookie 中的 BAIDUID
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, "BAIDUID") {
			baiduid = cookie.Value
			break
		}
	}
	if baiduid == "" {
		return "", "", fmt.Errorf("BAIDUID not found in Set-Cookie")
	}

	// 7. 读取响应 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading body: %w", err)
	}

	// 8. 提取 token
	re := regexp.MustCompile(`token:\s*'([a-zA-Z0-9]+)'`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", "", fmt.Errorf("token not found in body")
	}
	token = matches[1]

	// 返回 BAIDUID 和 token
	return baiduid, token, nil
}
