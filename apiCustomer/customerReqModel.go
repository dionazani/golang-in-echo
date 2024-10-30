package apiCustomer

type CustomerReqModel struct {
	Id                     int32  `json:"id"`
	CustomerName           string `json:"customeName"`
	CustomerGender         string `json:"customerGender"`
	CustomerIdentityNumber string `json:"customerIdentityNumber"`
}
