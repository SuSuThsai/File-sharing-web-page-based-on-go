package UpLoad

import (
	"github.com/jakeslee/aliyundrive"
	"log"
	"strconv"
	//第三方

	//目录
	"File-sharing-web-page/Global"
)

func IsExist() string {
	fullpath := ""
	if Global.Global.User.ID == 0 {
		fullpath = "/文件共享目录/未登录用户文件"
	} else {
		userid := "/" + strconv.Itoa(int(Global.Global.User.ID))
		secondPath := "/文件共享目录/用户文件列表"
		rsp0, _, err := Global.Global.Driver.ResolvePathToFileId(Global.Global.Credential, secondPath)
		if err != nil && err != aliyundrive.ErrPartialFoundPath {
			log.Println("寻找路径出错:", err.Error())
		}
		fullpath = secondPath + userid
		//fmt.Println(fullpath)
		rsp, _, err := Global.Global.Driver.ResolvePathToFileId(Global.Global.Credential, fullpath)
		//fmt.Println(rsp, rsp0)
		if rsp0 == rsp {
			Global.Global.Driver.CreateDirectory(Global.Global.Credential, rsp, strconv.Itoa(int(Global.Global.User.ID)))
			rsp, _, err = Global.Global.Driver.ResolvePathToFileId(Global.Global.Credential, fullpath)
			if err != nil && err != aliyundrive.ErrPartialFoundPath {
				log.Println("寻找路径出错:", err.Error())
			}
			//Global.Global.Driver.CreateDirectory(Global.Global.Credential,rsp,"图片")
			//Global.Global.Driver.CreateDirectory(Global.Global.Credential,rsp,"视频")
			//Global.Global.Driver.CreateDirectory(Global.Global.Credential,rsp,"音乐")
			Global.Global.Driver.CreateDirectory(Global.Global.Credential, rsp, "文件")
		} else if err != nil && err != aliyundrive.ErrPartialFoundPath {
			log.Println("寻找路径出错:", err.Error())
		}
	}
	return fullpath
}
