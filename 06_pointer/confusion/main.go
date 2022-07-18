package main

func main() {
	dd := &downloadFromDisk{
		secret:   &mobileTokenFynamic{mobileNumber: "123456"},
		filePath: "file.txt",
	}
	dd.downloadFile()
}

// 定义一个存手机号的结构体
type mobileTokenFynamic struct {
	mobileNumber string
}

// 给这个结构体添加 获取令牌 的方法
func (d *mobileTokenFynamic) GetSecret() string {
	return "something"
}

// 将获取令牌定义为接口
type DynamiSecret interface {
	GetSecret() string
}
