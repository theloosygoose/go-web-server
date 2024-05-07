package types

type Photo struct {
	ID          int
	Name        string
	Location    string
	Date        string
	Description string
	Image       ImageData
}

type ImageData struct {
	FileName string
	Height   int
	Width    int
}
