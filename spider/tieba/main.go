package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//爬取起始页，终止页
	var start, end int
	fmt.Println("请输入爬取起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入爬取终止页(>=start):")
	fmt.Scan(&end)
	work(start, end)
}

//爬取页面操作
func work(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)

	//循环爬取每一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i+1)*50)
		result, err := HttpGet(url)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("result=%s", result)

		//将读取到的网页保存到文件
		file, err := os.Create("第" + strconv.Itoa(i) + "页.html")
		if err != nil {
			fmt.Println("Create file err:", err)
			continue
		}
		file.WriteString(result)
		file.Close()
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	//循环读取网页数据
	bytes := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(bytes)
		if n == 0 {
			fmt.Println("网页读取完毕")
			break
		}
		if err != nil && err != io.EOF {
			return "", err
		}
		result += string(bytes[:n])
	}
	return result, nil
}
