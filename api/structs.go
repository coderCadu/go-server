package api

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Clean Architecture", Artist: "coderCadu", Price: 128.92},
	{ID: "2", Title: "Clean Design", Artist: "coderCadu", Price: 12.91},
	{ID: "3", Title: "Clean Code", Artist: "coderCadu", Price: 38.97},
}