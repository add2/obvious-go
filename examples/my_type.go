package examples

import (
	"fmt"
)

type some int

func (i some) value() string {
	return fmt.Sprintf("%#v", i)
}

func (i some) pointer() string {
	return fmt.Sprintf("%p", &i)
}

func (i some) _type() string {
	return fmt.Sprintf("%T", &i)
}
