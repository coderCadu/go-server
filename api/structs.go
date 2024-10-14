package api

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "João e Maria", Artist: "Chico Buarque", Price: 128.92},
	{ID: "2", Title: "Sunflower", Artist: "Post Malone and Swae Lee", Price: 12.91},
	{ID: "3", Title: "I'm yours", Artist: "Jason Mraz", Price: 38.97},
}