package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

//以数组的形式输出文件名, 会带相对路径
func dir() []string {
	files, _ := filepath.Glob("./*")
	return files
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func ls(p string) map[string]int64 {
	//确认目录
	files, _ := ioutil.ReadDir(p)
	//用一个map来储存当前目录的文件名和文件大小
	file := make(map[string] int64)
	for _, f := range files {
		//如果是文件夹，文件大小就为-1，到时候前端检测
		if f.IsDir() {
			file[f.Name()] = -1
		} else {
			file[f.Name()] = f.Size() //Decimal(float64(f.Size()) / 1048576)改在前端算了
		}
		//file[f.Name()] = f.IsDir() ? -1 : f.Size()
		//go不支持选择运算符真头大
	}
	return file
}

func main() {
	http.HandleFunc("/fileInfo", fileInfo)

	http.ListenAndServe(":8086", nil)
}

func fileInfo(writer http.ResponseWriter, request *http.Request) {
	httpCORS(writer,"*")
	files := ls("./")
	fmt.Println(files)
	filesJSON, err := json.Marshal(files)
	//filesJSON, err := json.MarshalIndent(files,"","\t")
	fmt.Println(err)
	fmt.Println(string(filesJSON))

	if request.Method == "GET" {
		fmt.Fprintf(writer, string(filesJSON))
	}
}

//跨域
func httpCORS(w http.ResponseWriter, url string) {
	w.Header().Set("Access-Control-Allow-Origin", url)
	w.Header().Add("Access-Control-Allow-Headers", "Access-Token, Content-Type")
	w.Header().Set("content-type", "application/json")
}