package request

import "encoding/json"

type JsonDataEntry struct {
	Table     string            `json:"table"`
	TableData []json.RawMessage `json:"tableData"`
}

type JsonData struct {
	Data []JsonDataEntry `json:"data"`
}
