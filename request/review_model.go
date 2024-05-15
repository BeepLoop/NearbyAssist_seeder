package request

type ReviewModel struct {
	ServiceId int `db:"serviceId"`
	Rating    int `db:"rating"`
}
