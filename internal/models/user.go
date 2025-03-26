package models

type User struct {
	ID        string `json:"user_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Phone     *string `json:"phone" bson:"phone"`
	Gender    *string `json:"gender" bson:"gender"` 
	BITS_Id    string `json:"bits_id" bson:"bits_id"`
}
