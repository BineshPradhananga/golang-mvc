package helpers

import (
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode int
	Status     int
	Message    string
	Data       interface{}
}

// func (h *Helper) GetData(data interface{}) {
// func GetData(c *fiber.Ctx, data interface{}) {
func GetData(data interface{}) {
	// return Response{
	// 	StatusCode: http.StatusOK,
	// 	Status:     200,
	// 	Message:    "OK",
	// 	Data:       data,
	// }
	var response Response
	response.StatusCode = http.StatusOK
	response.Status = 200
	response.Message = "OK"
	response.Data = data
	return string(&response)
	fmt.Println(response)
	// return response

	// return c.JSON(response.StatusCode, response)
	// return c.Status(200).JSON(json.response)
	// return &response
}
