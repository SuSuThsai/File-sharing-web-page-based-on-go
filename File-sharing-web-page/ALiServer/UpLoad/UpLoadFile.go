package UpLoad

import (
	"log"
	"path"
	//第三方
	"github.com/jakeslee/aliyundrive"
	"mime/multipart"
	//目录
	"File-sharing-web-page/ALiServer/Until"
	"File-sharing-web-page/Global"
)

func UpLoad(f *multipart.FileHeader) (fileId string) {
	fullpath := IsExist()
	f1, _ := f.Open()
	filetype := Until.UploadType(path.Ext(f.Filename))
	rsp, _, err := Global.Global.Driver.ResolvePathToFileId(Global.Global.Credential, fullpath+filetype)
	if err != nil && err != aliyundrive.ErrPartialFoundPath {
		log.Println("寻找路径出错:", err.Error())
	}
	file, err := Global.Global.Driver.UploadFile(Global.Global.Credential, &aliyundrive.UploadFileOptions{
		Name:             f.Filename,
		Size:             f.Size,
		ParentFileId:     rsp,
		ProgressStart:    nil,
		ProgressCallback: nil,
		ProgressDone:     nil,
		Reader:           f1,
	})
	if err != nil && err != aliyundrive.ErrPartialFoundPath {
		log.Println("上传文件失败:", err.Error())
	}
	if file == nil {
		fileId = ""
	} else {
		//fmt.Println(file.ParentFileId, file.Name)
		fileId = file.FileId
	}
	return
}
