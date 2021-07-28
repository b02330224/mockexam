package utils

import (
	"fmt"
	"os"
)

/*
主要用于输出错误与错误爆发的场景，并暴力退出程序（当然是在DEBUG模式下）
 */


/*处理错误：有错误时暴力退出*/
func HandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

