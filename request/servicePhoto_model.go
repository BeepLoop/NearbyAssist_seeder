package request

type ServicePhotoModel struct {
	VendorId  int    `db:"vendorId"`
	ServiceId int    `db:"serviceId"`
	Url       string `db:"url"`
}
