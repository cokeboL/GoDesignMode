package Action

import (
	"fmt"
)

type Singable struct {
}

type NotSingable struct {
}

func (obj *Singable) Sing() {
	fmt.Println("    singing ~~~")
}

func (obj *NotSingable) Sing() {
	fmt.Println("    can't sing !!!")
}
