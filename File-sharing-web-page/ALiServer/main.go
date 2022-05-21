package main

import (
	"File-sharing-web-page/ALiServer/ALiInits"
	"File-sharing-web-page/ALiServer/UpLoad"
	"fmt"
	"os"
)

//初始化
func main() {
	ALiInits.InitsConfig()
	ALiInits.InitsDriver()
	file, _ := os.Open("F:\\图片\\奈奈\\奈奈1.jpg")
	fileId := UpLoad.UpLoad(file)
	URLInfo := UpLoad.GetUrl(fileId)
	fmt.Println("", *URLInfo.Url, fileId)
}
