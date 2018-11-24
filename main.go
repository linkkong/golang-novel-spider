package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"spider/wx"
	"path/filepath"
)

const (
	Host       = "https://www.biqiuge.com"
)

type newStr struct {
	Title string
	Url   string
}

// gbk 转 utf8
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func main() {

	directory := getCurrentDirectory()

	msgToken := wx.GetToken()

	//圣墟
	var sxFileName = directory + "/sx.txt"
	var sxUrl = "https://www.biqiuge.com/book/4772/"
	check(sxFileName, sxUrl, msgToken)

	//飞剑问道
	var swordFileName = directory + "/fj.txt"
	var swordUrl = "https://www.biqiuge.com/book/24277/"
	check(swordFileName, swordUrl, msgToken)
}

//检查是否更新
func check(fileName string, uri string, token string) {
	cacheStr := readFile(fileName)

	res, err := http.Get(uri)
	if err != nil {
		fmt.Println("获取网址内容失败")
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var lastSix = ""
	var newStrs []newStr
	// Find the review items
	doc.Find(".listmain dd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//此处只截取前5个，即最新章节列表
		if i < 5 {
			aNode := s.Find("a")
			band := aNode.Text()
			url, _ := aNode.Attr("href")
			title := ConvertToString(band, "gbk", "utf-8")
			lastSix += title
			//判断当前title是否在文件缓存中，以此判断是否更新
			if !strings.Contains(cacheStr, title) {
				newStrs = append(newStrs, newStr{Title: title, Url: Host + url})
			}
		}
	})
	if len(newStrs) > 0 {
		writeFile(fileName, lastSix)
		log.Printf("%+v", newStrs)
		bookName := ""
		if strings.Contains(fileName, "sx.txt") {
			bookName = "圣墟"
		} else {
			bookName = "飞剑问道"
		}
		for k, v := range newStrs {
			if k < 2 {
				//sendMsg(bookName+"更新了"+v, token)
				wx.SendTemplateMsg(token, v.Url, bookName, v.Title)
			}
		}
	}
}

//写入文件
func writeFile(fileName string, writeString string) {
	var d1 = []byte(writeString)
	err2 := ioutil.WriteFile(fileName, d1, 0666) //写入文件(字节数组)
	if err2 != nil {
		log.Fatal(err2)
	}
}

//读取文件内容
func readFile(fileName string) string {
	exist, _ := PathExists(fileName)
	if !exist {
		return ""
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cacheStr := string(bytes)
	return cacheStr
}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取当前程序运行目录
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
