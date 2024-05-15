package request

type ServiceModel struct {
	VendorId    int      `db:"vendorId"`
	Description string   `db:"description"`
	Rate        string   `db:"rate"`
	Tags        []string `db:"tags"`
	Latitude    float64  `db:"latitude"`
	Longitude   float64  `db:"longitude"`
}
