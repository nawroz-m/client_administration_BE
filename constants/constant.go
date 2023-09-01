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

type UserLoginLocalStorage struct {
	Id  string             `bson:"user_id,omitempty" `
	Email  string             `bson:"email,omitempty" `
	Password  string             `bson:"password,omitempty" `
	Role  string             `bson:"role,omitempty" `
}

type Address struct {
    Street     string  `bson:"street,omitempty" `
    PostalCode int64  `bson:"postalcode,omitempty"`
    City       string  `bson:"city,omitempty" `
    Country    string  `bson:"country,omitempty" `
}

type UserInfoToUpdate struct {
	Id  string             `bson:"id,omitempty" `
	FirstName  string             `bson:"firstname,omitempty" `
	LastName  string             `bson:"lastname,omitempty" `
	Email  string             `bson:"email,omitempty" `
	Telephone  int64             `bson:"telephone,omitempty" `
	PostalAddress  Address        `bson:"postaladdress,omitempty" `
}

type ActiveDeactive struct {
	Id string					`bson:"id,omitempty"`
}



type SkipLimit struct {
	Skip int64		`bson:"skip,omitempty"`
	Limit int64		`bson:"limit,omitempty"`
}
type SearchUserData struct {
	Id string					`bson:"id,omitempty"`
	Search string				`bson:"search,omitempty"`
	Active bool					`bson:"active,omitempty"`
	ExtraParams SkipLimit		`bson:"extraParams,omitempty"`
	Count int64		`bson:"count,omitempty"`

}

type SearhResponse struct {
        Error   bool         `json:"error"`
        Message string       `json:"message"`
        Count   int64        `json:"count"`
        Data    interface{} `json:"data"`
	
}