package constants


type UserRegisterResponse struct{
	Email string `bson:"email,omitempty"`
	Id string   `bson:"id,omitempty"`
	Message string  `bson:"email,omitempty"`
}

type UserLoginInfo struct {
	Email  string             `bson:"email,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required"`	
}


type UserLoginPayload struct {
	Id  string             `bson:"user_id,omitempty" validate:"required"`
	Email  string             `bson:"email,omitempty" validate:"required"`
	Password  string             `bson:"password,omitempty" validate:"required"`
	Role  string             `bson:"role,omitempty" validate:"required"`
	Token  string             `bson:"token,omitempty" validate:"required"`
}

type Address struct {
    Street     string  `bson:"street,omitempty" `
    PostalCode int64  `bson:"postalcode,omitempty"`
    City       string  `bson:"city,omitempty" `
    Country    string  `bson:"country,omitempty" `
}

type UserInfoToUpdate struct {
	FirstName  string             `bson:"firstname,omitempty" `
	LastName  string             `bson:"lastname,omitempty" `
	Email  string             `bson:"email,omitempty" `
	telephone  int64             `bson:"telephone,omitempty" `
	PostalAddress  Address        `bson:"postaladdress,omitempty" `
}
