package next

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func routerHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()     //获取URL参数
	dir, _ := os.Getwd()           //获取当前文件所在目录
	urlPath := r.URL.Path          //获取URL文件相对路径
	file := dir + urlPath          //获取URL文件相对应的地址
	fileInfo, err := os.Stat(file) //获取文件状态
	if err == nil {
		if !fileInfo.IsDir() { //判断是否为目录
			http.ServeFile(w, r, file) //静态文件服务
			if len(urlParams) != 0 {
				fmt.Println(urlParams)
			}
		} else {
			for k, v := range urlParams {
				fmt.Println(k, v)
			}

		}
	}
}

func Run() {
	webRoot, _ := os.Getwd() //获取当前文件所在目录
	fmt.Println("WebRoot: " + webRoot + "\nWebUrl: http://localhost:9999\nListenAndServe Start\nBy ivup.cn")
	http.HandleFunc("/", routerHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
