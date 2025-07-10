package music

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetMusicUrl(ctx context.Context, id, source string) (murl *MusicUrl) {
	// 创建请求参数
	data := url.Values{}
	data.Set("types", "url")
	data.Set("id", id)
	data.Set("source", source)

	// 构建请求URL
	baseURL := "http://music.lessh.cn/api.php"
	callback := "jQuery111307641341149405615_1752153153635"
	fullURL := fmt.Sprintf("%s?callback=%s", baseURL, callback)

	// 创建HTTP POST请求
	req, err := http.NewRequest("POST", fullURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "http://music.lessh.cn")
	req.Header.Set("Referer", "http://music.lessh.cn/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36 Edg/138.0.0.0")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// 忽略SSL验证（对应curl的--insecure选项）
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true, // 对应curl默认行为
		},
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 输出原始响应（调试用）
	fmt.Println("原始响应内容:", string(body))

	// 处理JSONP响应，提取JSON部分
	murl = &MusicUrl{}
	json.Unmarshal(WashMusicJsonRespBody(body), murl)
	return murl
}
