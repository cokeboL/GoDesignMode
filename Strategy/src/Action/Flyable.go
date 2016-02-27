package Action

import (
	"fmt"
)

type Flyable struct {
}

func (obj *Flyable) Fly() {
	fmt.Println("    flying ~~~")
}

type NotFlyable struct {
}

func (obj *NotFlyable) Fly() {
	fmt.Println("    can't fly !!!")
}
