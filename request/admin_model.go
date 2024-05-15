package request

type AdminModel struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
