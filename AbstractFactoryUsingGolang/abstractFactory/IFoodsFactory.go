package abstractFactory

import (
	"fmt"

)

type IFoodsFactory interface {
	MakePizza() IPizza
	MakeSandwich() ISandwich
}

func GetFoodsFactory(typ string) (IFoodsFactory, error) {
	if typ == "veggy" {
		return &veggy{}, nil
	}

	if typ == "hawaiian" {
		return &hawaiian{}, nil
	}

	return nil, fmt.Errorf("Wrong type passed")
}
