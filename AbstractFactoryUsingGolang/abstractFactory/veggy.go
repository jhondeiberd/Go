package abstractFactory

type veggy struct {
}

func (a *veggy) MakePizza() IPizza {
	return &veggyPizza{
		pizza: pizza{
			typ: "veggy pizza",
			dough: "veggy dough",
			sauce: "veggy sauce",
			topping:"veggy topping",
		},
	}
}

func (a *veggy) MakeSandwich() ISandwich {
	return &veggySandwich{
		sandwich: sandwich{
			typ: "veggy sandwich",
			bread: "veggy bread",
			salade: "veggy salade",
			topping:"veggy topping",
		},
	}
}

