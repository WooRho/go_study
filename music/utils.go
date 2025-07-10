package music

import "strings"

func WashMusicJsonRespBody(body []byte) []byte {
	// 1. 去除JSONP包裹层，提取JSON部分
	respBody := string(body)
	start := strings.Index(respBody, "(") + 1
	end := strings.LastIndex(respBody, ")")
	jsonStr := respBody[start:end]
	return []byte(jsonStr)
}
