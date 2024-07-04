package request

type AdminModel struct {
    Username     string `json:"username" db:"username"`
    Password     string `json:"password" db:"password"`
    Role         string `json:"role" db:"role"`
    UsernameHash string `json:"usernameHash" db:"usernameHash"`
}
