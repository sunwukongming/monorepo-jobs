package bolejiang

type ComposedCity struct {
	DataCity
	Children []DataDistrict `json:"children"`
}
