package user

// Profile declare user information
type Profile struct {
	ID        int64  `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	Token     string `json:"token" db:"token"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}
