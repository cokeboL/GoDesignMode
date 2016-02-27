package Duck

import (
	"Action"
	"fmt"
)

type WildDuck struct {
	DuckAttr
	Action.Flyable
	Action.NotSingable
}

func (duck *WildDuck) Description() {
	duck.Age = 2
	duck.Weight = 3.5
	fmt.Printf("This is a wild duck, %d years old, weight is: %.2f\n", duck.Age, duck.Weight)
}
