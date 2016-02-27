package Duck

import (
	"Action"
)

type Duck interface {
	Action.Action
	Description()
}

type DuckAttr struct {
	Age    int
	Weight float32
}
