package request

type UserModel struct {
	Name      string `db:"name"`
	Email     string `db:"email"`
	ImageUrl  string `db:"imageUrl"`
	EmailHash string `db:"emailHash"`
}
