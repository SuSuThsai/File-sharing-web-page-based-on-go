package main

import (
	"fmt"
	"github.com/jakeslee/aliyundrive"
	"log"
	"os"
)

func main() {
	drive := aliyundrive.NewClient(&aliyundrive.Options{
		AutoRefresh: true,
		UploadRate:  8 * 1024 * 1024, // 限速 8MBps
	})
	fmt.Println(drive)
	cred, err := drive.AddCredential(aliyundrive.NewCredential(&aliyundrive.Credential{
		RefreshToken: "a1156232972e43d392461b1c8b09d772",
	}))
	//drive.RefreshAllToken()
	fmt.Println(cred)
	if err != nil {
		log.Panic("获取个人信息失败", err.Error())
	}
	filepath := "F:\\图片\\奈奈\\奈奈1.jpg"
	stat, _ := os.Stat(filepath)
	localfile, _ := os.Open(filepath)
	x := localfile.Name()
	fmt.Println(x)
	fullpath := "/文件共享目录/未登录用户文件/图片"
	rsp, _, _ := drive.ResolvePathToFileId(cred, fullpath)
	file, _ := drive.UploadFile(cred, &aliyundrive.UploadFileOptions{
		Name:             stat.Name(),
		Size:             stat.Size(),
		ParentFileId:     rsp,
		ProgressStart:    nil,
		ProgressCallback: nil,
		ProgressDone:     nil,
		Reader:           localfile,
	})
	fmt.Println(file.ParentFileId)
}
