package request

type ServiceModel struct {
	// VendorId    int        `db:"vendorId"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Rate        string     `db:"rate"`
	Tags        []TagModel `db:"tags"`
	Latitude    float64    `db:"latitude"`
	Longitude   float64    `db:"longitude"`
}
