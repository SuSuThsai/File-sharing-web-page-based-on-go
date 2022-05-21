package UpLoad

import (
	//本地
	"fmt"
	"log"
	"os"
	"path"
	//第三方
	"github.com/jakeslee/aliyundrive"
	//目录
	"File-sharing-web-page/ALiServer/ALiInits"
	"File-sharing-web-page/Global"
)

func UpLoad(f *os.File) (fileId string) {
	fullpath := IsExist()
	stat, _ := os.Stat(f.Name())
	filetype := ALiInits.UploadType(path.Ext(stat.Name()))
	rsp, _, err := Global.Global.Driver.ResolvePathToFileId(Global.Global.Credential, fullpath+filetype)
	if err != nil && err != aliyundrive.ErrPartialFoundPath {
		fmt.Println("寻找路径出错:", err.Error())
		log.Panic(err)
	}
	file, err := Global.Global.Driver.UploadFile(Global.Global.Credential, &aliyundrive.UploadFileOptions{
		Name:             stat.Name(),
		Size:             stat.Size(),
		ParentFileId:     rsp,
		ProgressStart:    nil,
		ProgressCallback: nil,
		ProgressDone:     nil,
		Reader:           f,
	})
	if err != nil && err != aliyundrive.ErrPartialFoundPath {
		fmt.Println("上传文件失败:", err.Error())
		log.Panic(err)
	}
	if file == nil {
		fileId = ""
	} else {
		fmt.Println(file.ParentFileId, file.Name)
		fileId = file.FileId
	}
	return
}
