package PizzaBuilder

type iBuilderPizza interface {
	setDough()
	setSauce()
	setTopping()
	getPizza() Pizza
}

func GetBuilder(builderType string) iBuilderPizza {
	if builderType == "veggy" {
		return &VeggyBuilderPizza{}
	}

	if builderType == "hawaiian" {
		return &HawaiianBuilderPizza{}
	}
	return nil
}
