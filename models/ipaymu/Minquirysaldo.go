package models

type InquirysaldoipayReq struct {
	Va string `json:"va" validate:"required"`
}

type InquirysaldoipayRes struct {
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
	Data    struct {
		Va              string `json:"Va"`
		MerchantBalance int    `json:"MerchantBalance"`
		MemberBalance   int    `json:"MemberBalance"`
	} `json:"Data"`
}
