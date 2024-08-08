package models

type User struct {
	ID       string   `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type AuthenticatedUser struct{
	ID       string   `json:"user_id" bson:"user_id"`
	Email    string   `json:"email" bson:"email"`
	Role     string   `json:"role" bson:"role"`
}
