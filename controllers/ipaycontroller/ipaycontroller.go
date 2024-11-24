package ipaycontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"eko/api-pg-bpr/helper"

	ipaymu "eko/api-pg-bpr/models/ipaymu"
)

var ResponseSuccessIpay = helper.ResponseSuccessIpay
var ResponseFailedIpay = helper.ResponseFailedIpay
var PaginateResponseSuccess = helper.PaginateResponseSuccess
var PaginateResponseFailed = helper.PaginateResponseFailed

func InquiryBalance(w http.ResponseWriter, r *http.Request) {
	var inquirybalance ipaymu.InquirysaldoipayReq // type untuk menangkap request dari klien
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&inquirybalance); err != nil {
		ResponseFailedIpay(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Membuat JSON body untuk request ke iPaymu
	requestBody, err := json.Marshal(map[string]string{
		"account": inquirybalance.Va,
	})
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// buat signature
	signature, timestamp, va := helper.GenerateSignature(requestBody, "POST")

	// Membuat request ke endpoint iPaymu
	ipayUrl := os.Getenv("IPAY_URL") + "/api/v2/balance"
	req, err := http.NewRequest("POST", ipayUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Menambahkan header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("signature", signature)
	req.Header.Set("va", va)
	req.Header.Set("timestamp", string(timestamp))

	// Mengirimkan request ke iPaymu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	// Membaca response dari iPaymu
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parsing response JSON ke model InquirysaldoipayRes
	var ipayResponse ipaymu.InquirysaldoipayRes
	if err := json.Unmarshal(body, &ipayResponse); err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, "Invalid response format")
		return
	}
	if !ipayResponse.Success {
		ResponseFailedIpay(w, ipayResponse.Status, ipayResponse.Message)
		return
	}

	// // Mengembalikan hanya field `Data` ke klien
	ResponseSuccessIpay(w, http.StatusOK, ipayResponse.Data)
}
func HistoryTransaction(w http.ResponseWriter, r *http.Request) {

	var historyTransactionReq ipaymu.HistoryTransactionReq // type untuk menangkap request dari klien
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&historyTransactionReq); err != nil {
		ResponseFailedIpay(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Membuat map JSON body secara dinamis
	jsonBody := make(map[string]interface{})

	// Properti lainnya ditambahkan langsung
	jsonBody["page"] = historyTransactionReq.Page
	jsonBody["orderBy"] = historyTransactionReq.OrderBy
	jsonBody["order"] = historyTransactionReq.Order
	jsonBody["limit"] = historyTransactionReq.Limit

	if historyTransactionReq.Account != "" {
		jsonBody["account"] = historyTransactionReq.Account
	}
	if historyTransactionReq.Id != 0 {
		jsonBody["id"] = historyTransactionReq.Id
	}
	if historyTransactionReq.Type != 0 {
		jsonBody["type"] = historyTransactionReq.Type
	}
	if historyTransactionReq.Status != 0 {
		jsonBody["status"] = historyTransactionReq.Status
	}
	if historyTransactionReq.BulkId != 0 {
		jsonBody["bulkId"] = historyTransactionReq.BulkId
	}
	if historyTransactionReq.Lang != "" {
		jsonBody["lang"] = historyTransactionReq.Lang
	}
	if historyTransactionReq.LockStatus != 0 {
		jsonBody["lockStatus"] = historyTransactionReq.LockStatus
	}
	if historyTransactionReq.Date != "" {
		jsonBody["date"] = historyTransactionReq.Date
	}
	if historyTransactionReq.Startdate != "" {
		jsonBody["startdate"] = historyTransactionReq.Startdate
	}
	if historyTransactionReq.Enddate != "" {
		jsonBody["enddate"] = historyTransactionReq.Enddate
	}

	// Marshal map ke JSON
	requestBody, err := json.Marshal(jsonBody)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(string(requestBody))

	// buat signature
	signature, timestamp, va := helper.GenerateSignature(requestBody, "POST")

	// Membuat request ke endpoint iPaymu
	ipayUrl := os.Getenv("IPAY_URL") + "/api/v2/history"
	req, err := http.NewRequest("POST", ipayUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Menambahkan header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("signature", signature)
	req.Header.Set("va", va)
	req.Header.Set("timestamp", string(timestamp))

	// Mengirimkan request ke iPaymu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	// Membaca response dari iPaymu
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(string(body))
	// Parsing response JSON ke model InquirysaldoipayRes
	var ipayResponse ipaymu.HistoryTransactionRes
	if err := json.Unmarshal(body, &ipayResponse); err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, "Invalid response format")
		return
	}
	if !ipayResponse.Success {
		ResponseFailedIpay(w, ipayResponse.Status, ipayResponse.Message)
		return
	}

	// Mengembalikan hanya field `Data` ke klien
	ResponseSuccessIpay(w, http.StatusOK, ipayResponse.Data)
}
func CheckTransaction(w http.ResponseWriter, r *http.Request) {
	var checkTrx ipaymu.CheckTransactionReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&checkTrx); err != nil {
		ResponseFailedIpay(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Membuat JSON body untuk request ke iPaymu
	requestBody, err := json.Marshal(map[string]interface{}{
		"account":       checkTrx.Account,
		"transactionId": checkTrx.TransactionId,
	})
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// buat signature
	signature, timestamp, va := helper.GenerateSignature(requestBody, "POST")

	// Membuat request ke endpoint iPaymu
	ipayUrl := os.Getenv("IPAY_URL") + "/api/v2/transaction"
	req, err := http.NewRequest("POST", ipayUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Menambahkan header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("signature", signature)
	req.Header.Set("va", va)
	req.Header.Set("timestamp", string(timestamp))

	// Mengirimkan request ke iPaymu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	// Membaca response dari iPaymu
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parsing response JSON ke model InquirysaldoipayRes
	var ipayCheckTrxResponse ipaymu.CheckTransactionRes
	if err := json.Unmarshal(body, &ipayCheckTrxResponse); err != nil {
		ResponseFailedIpay(w, http.StatusInternalServerError, "Invalid response format")
		return
	}

	if !ipayCheckTrxResponse.Success {
		ResponseFailedIpay(w, ipayCheckTrxResponse.Status, ipayCheckTrxResponse.Message)
		return
	}

	// Mengembalikan hanya field `Data` ke klien
	ResponseSuccessIpay(w, http.StatusOK, ipayCheckTrxResponse.Data)
}
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTransaction")
}
