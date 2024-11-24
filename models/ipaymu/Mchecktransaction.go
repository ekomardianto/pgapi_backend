package models

type CheckTransactionReq struct {
	TransactionId int    `json:"transactionId"`
	Account       string `json:"account"`
}

type CheckTransactionRes struct {
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
	Data    struct {
		TransactionId  int     `json:"TransactionId"`
		SessionId      string  `json:"SessionId"`
		ReferenceId    string  `json:"ReferenceId"`
		RelatedId      *string `json:"RelatedId"`
		Sender         string  `json:"Sender"`
		Receiver       string  `json:"Receiver"`
		SubTotal       int     `json:"SubTotal"`
		Amount         int     `json:"Amount"`
		Fee            int     `json:"Fee"`
		Status         int     `json:"Status"`
		StatusDesc     string  `json:"StatusDesc"`
		PaidStatus     string  `json:"PaidStatus"`
		IsLocked       bool    `json:"IsLocked"`
		Type           int     `json:"Type"`
		TypeDesc       string  `json:"TypeDesc"`
		Notes          *string `json:"Notes"`
		IsEscrow       bool    `json:"IsEscrow"`
		CreatedDate    string  `json:"CreatedDate"`
		ExpiredDate    string  `json:"ExpiredDate"`
		SuccessDate    string  `json:"SuccessDate"`
		SettlementDate string  `json:"SettlementDate"`
		PaymentMethod  string  `json:"PaymentMethod"`
		PaymentChannel string  `json:"PaymentChannel"`
		PaymentCode    string  `json:"PaymentCode"`
		PaymentName    string  `json:"PaymentName"`
		BuyerName      string  `json:"BuyerName"`
		BuyerPhone     string  `json:"BuyerPhone"`
		BuyerEmail     string  `json:"BuyerEmail"`
	} `json:"Data"`
}
