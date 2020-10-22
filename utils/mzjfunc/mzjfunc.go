package mzjfunc

import (
	"fmt"
	"runtime"
)

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	test()
}
func test() {
	fmt.Println("FuncName1 =", RunFuncName())
}
