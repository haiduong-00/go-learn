package main

import (
	"errors"
	"fmt"
	"time"
)

// Trong Go, trong ngôn ngữ khác: có thể tạo ra một error/exception riêng tùy ý
// Trong Go: các cách tạo error:
// fmt.Errorf(format string, a ...interface{}) error
// errors.New(text string) error
// Tạo 1 struct có method Error() string

type err struct{
	mainerr error
	when time.Time
}

func (t err) Error() string  {
	return fmt.Sprintf("An error name %v has occured at %v",t.mainerr,t.when)
}

func main()  {
	err := err{
		mainerr: errors.New("Lỗi"),
		when: time.Now(),
	}
	panic(err)
}