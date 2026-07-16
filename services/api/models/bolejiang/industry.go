package bolejiang

type ComposedIndustry struct {
	DataIndustry
	Children []ComposedIndustry
}
