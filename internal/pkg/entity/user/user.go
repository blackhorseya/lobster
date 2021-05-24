package user

// Profile declare a user basic information
type Profile struct {
	// ID describe the unique identify code of user
	ID int64 `json:"-" db:"sn"`

	// AccessToken describe this user's accessToken
	AccessToken string `json:"access_token" db:"access_token"`

	// Password describe user's password to login system
	Password string `json:"-" db:"password"`

	// Email describe user's email to login system
	Email string `json:"email" db:"email"`
}
