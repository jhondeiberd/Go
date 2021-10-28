package abstractFactory



type hawaiian struct {
}

func (n *hawaiian) MakePizza() IPizza {
	return &hawaiianPizza{
		pizza: pizza{
			typ: "hawaiian pizza",
			dough: "hawaiian dough",
			sauce: "hawaiian sauce",
			topping:"hawaiian topping",
		},
	}
}

func (n *hawaiian) MakeSandwich() ISandwich {
	return &hawaiianSandwich{
		sandwich: sandwich{
			typ: "hawaiian sandwich",
			bread: "hawaiian bread",
			salade: "hawaiian salade",
			topping:"hawaiian topping",
		},
	}
}

