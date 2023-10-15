package helpers

import (
	"net/http"
)

type Helper struct {
	status     int
	message    string
	statusCode int
	data       interface{}
}

// func (h *Helper) SuccessMessage(message string) interface{} {
func SuccessMessage(message string) interface{} {

	// var h = new(Helper)
	// h.status = 1
	// h.message = message
	// h.statusCode = http.StatusOK
	// return h

	response := map[string]interface{}{
		"status":     true,
		"statusCode": http.StatusOK,
		"message":    message,
	}
	return response

}

// func (h *Helper) GetData(data interface{}) {
// func GetData(c *fiber.Ctx, data interface{}) {
func GetData(data interface{}) interface{} {
	response := map[string]interface{}{
		"status":     true,
		"statusCode": http.StatusOK,
		"message":    "Data Found!!",
		"response": map[string]interface{}{
			"data": data,
		},
	}
	return response

}
