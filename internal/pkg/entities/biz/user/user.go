package user

// Profile declare a user basic information
type Profile struct {
	// ID describe the unique identify code of user
	ID string `json:"-" db:"sn"`

	// Password describe user's password to login system
	Password string `json:"-" db:"password"`

	// Email describe user's email to login system
	Email string `json:"email" db:"email"`

	// SignupAt describe user signup platform milliseconds
	SignupAt int64 `json:"-" db:"signupAt"`
}
