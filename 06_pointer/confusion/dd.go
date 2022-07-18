package main

import "fmt"

type downloadFromDisk struct {
	// 令牌字段 为获取令牌定义的接口
	secret   DynamiSecret
	filePath string
}

func (dd *downloadFromDisk) downloadFile() {
	if err := dd.logincheck(); err != nil {
		// todo 重新登录
	}
	dd.downloadFromAliYun(dd.filePath)
	fmt.Println("sueeccd")
}

func (dd *downloadFromDisk) logincheck() error {
	dd.checkSecret(dd.secret.GetSecret())
	return nil
}

func (dd *downloadFromDisk) checkSecret(secret string) {

}

func (dd *downloadFromDisk) downloadFromAliYun(file string) {

}
