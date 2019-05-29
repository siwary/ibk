package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 지정일자로, 표시포맷 설정 가능
	now := time.Now().Format("2006-01-02, 15:04:05")
	p(now)
}
