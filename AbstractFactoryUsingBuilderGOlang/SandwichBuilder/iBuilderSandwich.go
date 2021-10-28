package SandwichBuilder

type iBuilderSandwich interface {
	setBread()
	setSalade()
	setTopping()
	getSandwich() Sandwich
}

func GetBuilder(builderType string) iBuilderSandwich {
	if builderType == "veggy" {
		return &VeggyBuilderSandwich{}
	}

	if builderType == "hawaiian" {
		return &HawaiianBuilderSandwich{}
	}
	return nil
}
