package utils

import (
	"github.com/petersunbag/coven"
)

func NewConverter(dst, src interface{}) (*coven.Converter, error) {
	return coven.NewConverter(dst, src)
}

func Converter(dst, src interface{}) error {
	c, err := coven.NewConverter(dst, src)
	if err != nil {
		return err
	}
	c.Convert(dst, src)
	return nil
}
