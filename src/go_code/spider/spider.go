package spider

import (
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
	"regexp"
	"time"
	_ "fmt"
)
var headers map[string]string = map[string]string{
	"User-Agent" : "golang",
	"Referer" : "",
} 

type Spider struct {
	Types string
	Url string
	Regex string
	Deep int64
	Client *http.Client
	Req *http.Request 
	ResponseBody []byte
	Headers map[string]string
}

func SetHeaders(header map[string]string) {
	for k, v := range header{
		header[k] = v
	}
}

func (sp *Spider) setHeader(header map[string]string) {
	for k, v := range header{
		sp.Req.Header.Add(k, v)
	}
}
func (sp *Spider) getContent() {
	sp.Client = &http.Client{}
	sp.Req, _ = http.NewRequest(sp.Types, sp.Url, nil)
	sp.setHeader(headers)
	resp, err := sp.Client.Do(sp.Req)
	if err != nil {
		log.Fatal(err)
	}
	//请求完关闭
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	sp.ResponseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
func (sp *Spider) SetParam(param map[string]string) {
	for k, v := range param {
		switch k {
			case "Url":
				sp.Url = v
			case "Regex":
				sp.Regex = v
			case "Deep":
				Deep, _ := strconv.ParseInt(v, 10, 0)
				sp.Deep = Deep
			case "Types":
				sp.Types = v
			default:
				panic("undefined param for spider struct")
		}
	}
}

func (sp *Spider) Find() []string {
	sp.getContent()
	responseBody := string(sp.ResponseBody)
	if sp.Regex == "" {
		panic("undefined param for spider struct of regex!")	
	}
	findKey := regexp.MustCompile(sp.Regex)
	findResult := findKey.FindAllString(responseBody, -1)
	//去重
	if len(findResult) > 0 {
		var filter map[string]int
		var newResult []string
		filter = make(map[string]int)
		for _, v := range findResult {
			if filter[v] == 0 {
				filter[v] = 1
				// fmt.Println(v)
				newResult = append(newResult, v)
				// fmt.Println(newResult)
			} else if (filter[v] >= 1) {
				continue
			}
		}
		findResult = newResult
	}
	return findResult
}

func (sp *Spider) Download(fileNamePath string, fileTypes string, downloadUrl string, fileName string) string {
	//todo
	if fileName == "" {
		times := time.Now().UnixNano()
		fileName = strconv.FormatInt(times, 10) + "." + fileTypes
	}
	savePath := fileNamePath + fileName
	sp.Url = downloadUrl
	sp.Client = &http.Client{}
	sp.Req, _ = http.NewRequest(sp.Types, sp.Url, nil)
	sp.setHeader(headers)
	resp, err := sp.Client.Do(sp.Req)
	if err != nil {
		log.Fatal(err)
	}
	//请求完关闭
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	sp.ResponseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(savePath, sp.ResponseBody, 777)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return savePath
}

