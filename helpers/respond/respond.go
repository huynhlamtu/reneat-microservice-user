package respond

import "reneat-microservice-user/constant"

type Respond struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InternalServerError() Respond {
	return Respond{
		Code:    constant.INTERNAL_SERVER_ERROR,
		Message: "Internal error",
		Data:    nil,
	}
}
