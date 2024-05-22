package types

type Photo struct {
	ID          int
	Name        string
	Location    string
	Date        string
	Description string
	Image       ImageData
	Collection  Collection
}

type ImageData struct {
	FileName string
	Height   string
	Width    string
}

type Collection struct {
	ID     int
	Name   string
	Photos []Photo
}
