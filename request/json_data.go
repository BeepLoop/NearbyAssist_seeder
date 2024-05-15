package request

type JsonDataEntry struct {
	Table     string        `json:"table"`
	TableData []interface{} `json:"tableData"`
}

type JsonData struct {
	Data []JsonDataEntry `json:"data"`
}
