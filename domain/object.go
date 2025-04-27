package domain

type Note struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}

type User struct {
	Id       string `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
