package models

type HistoryTransactionReq struct {
	Date       string `json:"date"`
	Startdate  string `json:"startdate"`
	Enddate    string `json:"enddate"`
	Page       int    `json:"page"`
	OrderBy    string `json:"orderBy"`
	Order      string `json:"order"`
	Limit      int    `json:"limit"`
	Account    string `json:"account"`
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	BulkId     int    `json:"bulkId"`
	Lang       string `json:"lang"`
	LockStatus int    `json:"lockStatus"`
}

type HistoryTransactionRes struct {
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
	Data    struct {
		Transaction []struct {
			TransactionId  int     `json:"TransactionId"`
			SessionId      string  `json:"SessionId"`
			ReferenceId    string  `json:"ReferenceId"`
			RelatedId      *int    `json:"RelatedId"`
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
		} `json:"Transaction"`
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"Pagination"`
	} `json:"Data"`
}
