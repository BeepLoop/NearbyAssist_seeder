package request

type ServiceTagModel struct {
	ServiceId int    `db:"serviceId"`
	TagTitle  string `db:"tagTitle"`
}
