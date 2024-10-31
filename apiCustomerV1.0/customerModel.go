package apiCustomer

type CustomerModel struct {
	Id                     int32  `json:"id"`
	CustomerName           string `json:"customerName"`
	CustomerGender         string `json:"customerGender"`
	CustomerIdentityNumber string `json:"customerIdentityNumber"`
	CustomerBirthPlace     string `json:"customerBirthPlace"`
	CustomerBirthDate      string `json:"customerBirthDate"`
}
