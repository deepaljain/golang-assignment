package movies

type Movie struct {
	Name   string  `json:"name"` 
	Actors []Actor `json:"actors"`
}

type Actor struct {
	Name string `json:"name"`
}
