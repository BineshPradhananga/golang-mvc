package helpers

import "net/http"

func SuccessMessage(message string) interface{} {
	response := map[string]interface{}{
		"status":     true,
		"statusCode": http.StatusOK,
		"message":    message,
	}
	return response

}

func ErrorMessage(message string) interface{} {

	response := map[string]interface{}{
		"status":     false,
		"statusCode": http.StatusForbidden,
		"message":    message,
	}
	return response

}

func UnauthorizedMessage(message string) interface{} {

	response := map[string]interface{}{
		"status":     false,
		"statusCode": http.StatusUnauthorized,
		"message":    message,
	}
	return response

}

func ValidationMessage(message string) interface{} {

	response := map[string]interface{}{
		"status":     false,
		"statusCode": http.StatusBadRequest,
		"message":    message,
	}
	return response

}

func GetData(data interface{}, message string) interface{} {
	if message == "" {
		message = "Data Found!!"
	}
	response := map[string]interface{}{
		"status":     true,
		"statusCode": http.StatusOK,
		"message":    message,
		"response": map[string]interface{}{
			"data": data,
		},
	}
	return response

}
