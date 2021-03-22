package dto

type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
