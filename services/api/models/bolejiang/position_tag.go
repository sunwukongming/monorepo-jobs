package bolejiang

type ComposedPositionTag struct {
	DataPositionTag
	Children []ComposedPositionTag
}
