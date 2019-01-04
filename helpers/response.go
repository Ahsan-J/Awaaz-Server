package helpers

import (
	"net/http"
	"fmt"
	"encoding/json"
)

// Response Struct to send
type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Message string	`json:"message"`
	Code string	`json:"code"`
}

// SendAPI sends resposne with given data
func (res *Response) SendAPI(w http.ResponseWriter,data interface{}) {
	// Writing data to the response method
	res.Data = data;
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err);
	}
}