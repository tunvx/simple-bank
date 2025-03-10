package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	// authURL            = "http://auth.banking.local"
	// cusmanURL          = "http://cusman.banking.local"
	authURL               = "http://127.0.0.1:8081"
	cusmanURL             = "http://127.0.0.1:8082"
	numUsers              = 50000 // Total number of users to create
	maxConcurrentRequests = 20    // Maximum concurrent requests
)

type Customer struct {
	CustomerRid      string `json:"customerRid"`
	FullName         string `json:"fullName"`
	DateOfBirth      string `json:"dateOfBirth"`
	PermanentAddress string `json:"permanentAddress"`
	PhoneNumber      string `json:"phoneNumber"`
	EmailAddress     string `json:"emailAddress"`
	CustomerTier     string `json:"customerTier"`
	CustomerSegment  string `json:"customerSegment"`
	FinancialStatus  string `json:"financialStatus"`
}

type Credentials struct {
	CustomerRid string `json:"customerRid"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	AccountNumber string `json:"accountNumber"`
	CurrencyType  string `json:"currencyType"`
}

func sendPostRequest(url string, data interface{}, token string) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("❌ Request failed [%d] - %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

func createUser(userID int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}        // Acquire semaphore
	defer func() { <-sem }() // Release semaphore

	userIDStr := fmt.Sprintf("%05d", userID)
	customerData := Customer{
		CustomerRid:      "000000" + userIDStr,
		FullName:         "Nguyen Van ABC",
		DateOfBirth:      "2001/01/01",
		PermanentAddress: "Tu Xa, Lam Thao, Phu Tho",
		PhoneNumber:      "+84000" + userIDStr,
		EmailAddress:     "example" + userIDStr + "@gmail.com",
		CustomerTier:     "gold",
		CustomerSegment:  "retail",
		FinancialStatus:  "good",
	}
	_, err := sendPostRequest(cusmanURL+"/v1/customers", customerData, "")
	if err != nil {
		fmt.Println("❌ Failed to create customer", userIDStr, "-", err)
		return
	}

	credentialsData := Credentials{
		CustomerRid: "000000" + userIDStr,
		Username:    "cus" + userIDStr,
		Password:    "App-0000",
	}
	_, err = sendPostRequest(authURL+"/v1/customers/credentials", credentialsData, "")
	if err != nil {
		fmt.Println("❌ Failed to create credentials", userIDStr, "-", err)
		return
	}

	loginData := Login{
		Username: "cus" + userIDStr,
		Password: "App-0000",
	}
	resp, err := sendPostRequest(authURL+"/v1/customers/login", loginData, "")
	if err != nil {
		fmt.Println("❌ Failed to login for user", userIDStr, "-", err)
		return
	}

	var loginResp map[string]interface{}
	json.Unmarshal(resp, &loginResp)
	accessToken, exists := loginResp["access_token"].(string)
	if !exists {
		fmt.Println("❌ Failed to get access token for user", userIDStr)
		return
	}

	accountData := Account{
		AccountNumber: "000000" + userIDStr,
		CurrencyType:  "VND",
	}
	_, err = sendPostRequest(cusmanURL+"/v1/accounts", accountData, accessToken)
	if err != nil {
		fmt.Println("❌ Failed to create bank account for", userIDStr, "-", err)
		return
	}

	fmt.Println("✅ Successfully created user", userIDStr)
}

func createUsersParallel() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentRequests) // Semaphore to limit concurrency

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go createUser(i, &wg, sem)
	}

	wg.Wait()
}

func main() {
	createUsersParallel()
}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sync"
// )

// const (
// 	baseURL               = "http://127.0.0.1"
// 	authsPort             = "8081"
// 	mansPort              = "8082"
// 	numUsers              = 20000 // Total number of users to create
// 	maxConcurrentRequests = 40    // Maximum concurrent requests
// )

// type Customer struct {
// 	CustomerRid      string `json:"customerRid"`
// 	FullName         string `json:"fullName"`
// 	DateOfBirth      string `json:"dateOfBirth"`
// 	PermanentAddress string `json:"permanentAddress"`
// 	PhoneNumber      string `json:"phoneNumber"`
// 	EmailAddress     string `json:"emailAddress"`
// 	CustomerTier     string `json:"customerTier"`
// 	CustomerSegment  string `json:"customerSegment"`
// 	FinancialStatus  string `json:"financialStatus"`
// }

// type Credentials struct {
// 	CustomerRid string `json:"customerRid"`
// 	Username    string `json:"username"`
// 	Password    string `json:"password"`
// }

// type Login struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Account struct {
// 	AccountNumber string `json:"accountNumber"`
// 	CurrencyType  string `json:"currencyType"`
// }

// func sendPostRequest(url string, data interface{}, token string) ([]byte, error) {
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	if token != "" {
// 		req.Header.Set("Authorization", "Bearer "+token)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("❌ Request failed [%d] - %s", resp.StatusCode, string(body))
// 	}

// 	return io.ReadAll(resp.Body)
// }

// func createUser(userID int, wg *sync.WaitGroup, sem chan struct{}) {
// 	defer wg.Done()
// 	sem <- struct{}{}        // Acquire semaphore
// 	defer func() { <-sem }() // Release semaphore

// 	userIDStr := fmt.Sprintf("%05d", userID)
// 	customerData := Customer{
// 		CustomerRid:      "000000" + userIDStr,
// 		FullName:         "Nguyen Van ABC",
// 		DateOfBirth:      "2001/01/01",
// 		PermanentAddress: "Tu Xa, Lam Thao, Phu Tho",
// 		PhoneNumber:      "+84000" + userIDStr,
// 		EmailAddress:     "example" + userIDStr + "@gmail.com",
// 		CustomerTier:     "gold",
// 		CustomerSegment:  "retail",
// 		FinancialStatus:  "good",
// 	}
// 	_, err := sendPostRequest(baseURL+":"+mansPort+"/v1/customers", customerData, "")
// 	if err != nil {
// 		fmt.Println("❌ Failed to create customer", userIDStr, "-", err)
// 		return
// 	}

// 	credentialsData := Credentials{
// 		CustomerRid: "000000" + userIDStr,
// 		Username:    "cus" + userIDStr,
// 		Password:    "App-0000",
// 	}
// 	_, err = sendPostRequest(baseURL+":"+authsPort+"/v1/customers/credentials", credentialsData, "")
// 	if err != nil {
// 		fmt.Println("❌ Failed to create credentials", userIDStr, "-", err)
// 		return
// 	}

// 	loginData := Login{
// 		Username: "cus" + userIDStr,
// 		Password: "App-0000",
// 	}
// 	resp, err := sendPostRequest(baseURL+":"+authsPort+"/v1/customers/login", loginData, "")
// 	if err != nil {
// 		fmt.Println("❌ Failed to login for user", userIDStr, "-", err)
// 		return
// 	}

// 	var loginResp map[string]interface{}
// 	json.Unmarshal(resp, &loginResp)
// 	accessToken, exists := loginResp["access_token"].(string)
// 	if !exists {
// 		fmt.Println("❌ Failed to get access token for user", userIDStr)
// 		return
// 	}

// 	accountData := Account{
// 		AccountNumber: "000000" + userIDStr,
// 		CurrencyType:  "VND",
// 	}
// 	_, err = sendPostRequest(baseURL+":"+mansPort+"/v1/accounts", accountData, accessToken)
// 	if err != nil {
// 		fmt.Println("❌ Failed to create bank account for", userIDStr, "-", err)
// 		return
// 	}

// 	fmt.Println("✅ Successfully created user", userIDStr)
// }

// func createUsersParallel() {
// 	var wg sync.WaitGroup
// 	sem := make(chan struct{}, maxConcurrentRequests) // Semaphore to limit concurrency

// 	for i := 1; i <= numUsers; i++ {
// 		wg.Add(1)
// 		go createUser(i, &wg, sem)
// 	}

// 	wg.Wait()
// }

// func main() {
// 	createUsersParallel()
// }
