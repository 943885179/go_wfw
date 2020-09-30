package mzjcode

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)
func GetRandCode(width int)string{//获取随机数
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var str="%0"+strconv.Itoa(width)+"v"
	return fmt.Sprintf(str, rnd.Int31n(int32(math.Pow10(width))))
}
func Default()string  {
	return GetRandCode(6)
}
