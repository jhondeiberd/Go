package SandwichBuilder

type iBuilderSandwich interface {
	setBread()
	setSalade()
	setTopping()
	getSandwich() sandwich
}

func GetBuilder(builderType string) iBuilderSandwich {
	if builderType == "veggy" {
		return &veggyBuilderSandwich{}
	}

	if builderType == "hawaiian" {
		return &hawaiianBuilderSandwich{}
	}
	return nil
}