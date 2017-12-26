package model

//	"time"

type TestSK struct {
	ID          int64  `meddler:"testsk_id,pk"	json:"-"`
	TestTime    string `meddler:"testsk_time"	json:"testsk_time"`
	TestAddress string `meddler:"testsk_add"	json:"testsk_add"`
	TestName    string `meddler:"testsk_name"	json:"testsk_name"`
	TestCount   string `meddler:"testsk_count"	json:"testsk_count"`
}
