package music

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetMusic(ctx context.Context) {
	musicURL := "http://m701.music.126.net//20250710215748//d2747b7e225c92b9f6e1629c67c09743//jdymusic//obj//wo3DlMOGwrbDjj7DisKw//14096492024//f1e7//73b5//a1a7//f84ec4b1179f60f9dec23ae3adf0b761.mp3?vuutv=IX9pf/jFgY985DuVHsRRGV75jt98FRhB7wtiffkT4CBH8zjGVk+Da81PrbbU3BWUG3+fmg6fHTrgvms40ZwERWWjIoD4DV6GtOx2ynE1AnJA8NqDGRr+zIAvEWQq2mzO"

	// 创建保存目录
	saveDir := "./downloads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.MkdirAll(saveDir, 0755)
	}

	// 从URL中提取文件名
	fileName := extractFileName(musicURL)
	savePath := filepath.Join(saveDir, fileName)

	// 下载文件
	err := downloadFile(musicURL, savePath)
	if err != nil {
		fmt.Printf("下载失败: %v\n", err)
		return
	}

	fmt.Printf("音乐已成功下载到: %s\n", savePath)
}

// 从URL中提取文件名
func extractFileName(url string) string {
	// 处理查询参数之前的部分
	path := url
	if idx := strings.Index(url, "?"); idx != -1 {
		path = url[:idx]
	}

	// 提取最后一个路径段
	segments := strings.Split(path, "/")
	if len(segments) == 0 {
		return "unknown.mp3"
	}

	// 使用最后一个段作为文件名
	return segments[len(segments)-1]
}

// 下载文件到指定路径
func downloadFile(url, savePath string) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 设置必要的请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Referer", "http://music.lessh.cn/") // 根据实际情况设置Referer

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败: %s", resp.Status)
	}

	// 创建文件
	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 复制内容到文件
	_, err = io.Copy(out, resp.Body)
	return err
}
