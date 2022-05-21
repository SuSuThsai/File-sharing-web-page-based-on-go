package DeleteFile

import (
	"File-sharing-web-page/Global"
	"fmt"
	"log"
)

func DeleteFile(fileId string) (err error) {
	base, err := Global.Global.Driver.RemoveFile(Global.Global.Credential, fileId)
	if err != nil {
		fmt.Println("删除失败:", err.Error())
		log.Fatal(err.Error())
	}
	fmt.Println(base.Code, base.Message)
	return
}
