package PizzaBuilder

type iBuilderPizza interface {
	setDough()
	setSauce()
	setTopping()
	getPizza() pizza
}

func GetBuilder(builderType string) iBuilderPizza {
	if builderType == "veggy" {
		return &veggyBuilderPizza{}
	}

	if builderType == "hawaiian" {
		return &hawaiianBuilderPizza{}
	}
	return nil
}