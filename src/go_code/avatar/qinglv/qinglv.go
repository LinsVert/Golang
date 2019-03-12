package qinglv

import(
	lsp "go_code/spider"
	"fmt"
)

var Url string = "http://www.wxcha.com/touxiang/qinglv/hot_1.html"

//获取第一层的数据
func GetUrls() []string {
	fmt.Printf("开始抓去第一层数据...\n")
	var readySearch map[string]int
	var param map[string]string
	readySearch = make(map[string]int)
	//默认查询第一条数据
	readySearch[Url] = 1
	param = make(map[string]string)
	param["Url"] = Url
	param["Deep"] = "1"
	param["Regex"] = `http:\/\/www.wxcha.com\/touxiang\/qinglv\/hot_\d+.html`
	param["Types"] = "get"
	sp := &lsp.Spider{}
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Referer": Url,
	}
	lsp.SetHeaders(headers)
	sp.SetParam(param)
	resp := sp.Find()
	result := resp
	if len(resp) > 0 {
		//有数据才继续
		i := sp.Deep
		var flag int64 = 0
		maxFor := len(resp)
		forLen := 0
		for flag < i {
			flag++
			for len(resp) > 0 && forLen <= maxFor {
				target := resp[0]
				if readySearch[target] == 0 {
					readySearch[target] = 1
					sp.Url = target
					resp = resp[1 : len(resp)]
					_resp := sp.Find()
					resp = append(resp, _resp...)
					result = append(result, _resp...)
				} else {
					resp = resp[1 : len(resp)]
				}
				forLen++
			}
			forLen = 0 
			// for _,v := range resp {
			// 	if readySearch[v] == 0 {
			// 		//没有采集过的 继续采集
			// 		fmt.Printf("flag is %d result for url %s\n", flag, v)
			// 		sp.Url = v
			// 	} else {
			// 		continue
			// 	}
			// 	resp = sp.Find()
			// 	result = append(result, resp...)
			// }
		}
	}
	//去重 想下是否可以想个办法简化这个去重
	var finalResultMap map[string]int
	finalResultMap = make(map[string]int)
	for _, v := range result {
		finalResultMap[v] = 1
	}
	var finalResultSlice []string
	for k, _ := range finalResultMap {
		finalResultSlice = append(finalResultSlice, k)
		// fmt.Printf("result for url %s\n", k)
	}
	// fmt.Println(finalResultSlice)
	return finalResultSlice
}

func GetDetailUrls() []string {
	fmt.Printf("开始抓去第二层数据...\n")
	// var readySearch map[string]int
	var param map[string]string
	// readySearch = make(map[string]int)
	param = make(map[string]string)
	param["Deep"] = "1"
	param["Regex"] = `http://www.wxcha.com/touxiang/\d+.html`
	param["Types"] = "get"
	sp := &lsp.Spider{}
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Referer": Url,
	}
	lsp.SetHeaders(headers)
	sp.SetParam(param)
	firstUrls := GetUrls()
	var result []string
	for _,v := range firstUrls {
		sp.Url = v
		resp := sp.Find()
		result = append(result, resp...)
	}
	//去重 想下是否可以想个办法简化这个去重
	var finalResultMap map[string]int
	finalResultMap = make(map[string]int)
	for _, v := range result {
		finalResultMap[v] = 1
	}
	var finalResultSlice []string
	for k, _ := range finalResultMap {
		finalResultSlice = append(finalResultSlice, k)
		// fmt.Printf("result for url %s\n", k)
	}
	// fmt.Println(finalResultSlice)
	return finalResultSlice
}


func DownloadImg(){
	//todo 下载图片
	fmt.Printf("开始下载图片...\n")
	var downLoadFile string = "/Users/lins/Desktop/spider/"
	// var readySearch map[string]int
	var param map[string]string
	// readySearch = make(map[string]int)
	param = make(map[string]string)
	param["Regex"] = `http://img.wxcha.com/file/\d+/\d+/\w+.jpg`
	param["Types"] = "get"
	sp := &lsp.Spider{}
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		"Referer": Url,
	}
	lsp.SetHeaders(headers)
	sp.SetParam(param)
	var result []string
	detailUrls := GetDetailUrls()
	for _,v := range detailUrls {
		sp.Url = v
		resp := sp.Find()
		result = append(result, resp...)
	}
	//去重 想下是否可以想个办法简化这个去重
	var finalResultMap map[string]int
	finalResultMap = make(map[string]int)
	for _, v := range result {
		finalResultMap[v] = 1
	}
	var finalResultSlice []string
	for k, _ := range finalResultMap {
		downLoadPath := sp.Download(downLoadFile, "jpg", k, "")
		if downLoadPath == "" {
			continue
		}
		fmt.Printf("result %s\n", downLoadPath)
		finalResultSlice = append(finalResultSlice, downLoadPath)
	}
	

}