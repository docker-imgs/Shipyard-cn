package v1

import (
	"fmt"
)

/*
	错误信息类型
 */
type Error struct {
	StatusCode int  //状态代码
	Status     string //状态
	msg        string //消息内容
}

 /**
  打印错误信息
  */
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.msg)
}
