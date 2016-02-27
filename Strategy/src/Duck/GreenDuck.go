package Duck

import (
	"Action"
	"fmt"
)

type GreenDuck struct {
	DuckAttr
	Action.NotFlyable
	Action.Singable
}

func (duck *GreenDuck) Description() {
	duck.Age = 1
	duck.Weight = 4.2
	fmt.Printf("This is a green duck, %d years old, weight is: %.2f\n", duck.Age, duck.Weight)
}
