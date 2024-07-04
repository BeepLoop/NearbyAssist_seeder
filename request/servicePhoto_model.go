package request

type ServicePhotoModel struct {
	Vendor    string `json:"vendor" db:"vendor"`
	ServiceId int    `json:"serviceId" db:"serviceId"`
	Url       string `json:"url" db:"url"`
}
