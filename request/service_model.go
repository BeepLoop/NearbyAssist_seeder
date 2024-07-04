package request

type ServiceModel struct {
	Vendor      string              `json:"vendor" db:"vendor"`
	Description string              `json:"description" db:"description"`
	Rate        string              `json:"rate" db:"rate"`
	Tags        []TagModel          `json:"tags" db:"tags"`
	Latitude    float64             `json:"latitude" db:"latitude"`
	Longitude   float64             `json:"longitude" db:"longitude"`
	Images      []ServicePhotoModel `json:"images" db:"images"`
	Reviews     []ReviewModel       `json:"reviews" db:"reviews"`
}
