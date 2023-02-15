package database


type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}


var Users = []User{

	{ID: 1, Name: "Sunny", Surname: "Flower", Email: "sunny@fake.org", IsActive: true},
}


