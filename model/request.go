package model

type Request struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

type RequestBillpay struct {
	StartDate   string `query:"start_date"`
	EndDate     string `query:"end_date"`
	Page        int    `query:"page"`
	Limit       int    `query:"limit"`
	Status      string `query:"status"`
	ChannelID   int    `query:"channel_id"`
	ReferenceID string `query:"reference_id"`
	Field       string `query:"field"`
}
