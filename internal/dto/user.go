package dto

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}
