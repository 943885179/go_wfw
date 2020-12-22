package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Code struct {
	Width int //随机数长度
}

func (c *Code) Get() string { //获取随机数
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var str = "%0" + strconv.Itoa(c.Width) + "v"
	return fmt.Sprintf(str, rnd.Int31n(int32(math.Pow10(c.Width))))
}
