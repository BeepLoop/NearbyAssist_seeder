package request

type VendorModel struct {
	Email string `json:"email" db:"email"`
	Job   string `json:"job" db:"job"`
}
