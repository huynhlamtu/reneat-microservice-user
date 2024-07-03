package respond

import (
	"fmt"
	"reneat-microservice-user/constant"
)

func Success(data interface{}, message string) Respond {
	return Respond{
		Code:    constant.SUCCESS,
		Message: message,
		Data:    data,
	}
}

func MissingParams() Respond {
	return Respond{
		Code:    constant.MISSING_PARAMS,
		Message: "Missing params",
		Data:    nil,
	}
}

func CreatedFail() Respond {
	return Respond{
		Code:    constant.CREATED_FAIL,
		Message: "Created fail!",
		Data:    nil,
	}
}

func UpdatedFail() Respond {
	return Respond{
		Code:    constant.UPDATED_FAIL,
		Message: "Updated fail!",
		Data:    nil,
	}
}

func DeletedFail() Respond {
	return Respond{
		Code:    constant.DELETE_FAIL,
		Message: "Delete fail!",
		Data:    nil,
	}
}

func Unauthorized() Respond {
	return Respond{
		Code:    constant.UNAUTHORIZED,
		Message: "Unauthorized",
		Data:    nil,
	}
}

func Forbidden() Respond {
	return Respond{
		Code:    constant.FORBIDDEN,
		Message: "Forbidden",
		Data:    nil,
	}
}

func ManyRequest() Respond {
	return Respond{
		Code:    constant.MANY_REQUEST,
		Message: "Too many request",
		Data:    nil,
	}
}

func NotFound() Respond {
	return Respond{
		Code:    constant.NOT_FOUND,
		Message: "Not found",
		Data:    nil,
	}
}

func MissingHeader() Respond {
	return Respond{
		Code:    constant.MISSING_HEADER,
		Message: "Missing request header",
		Data:    nil,
	}
}

func InValidParams() Respond {
	return Respond{
		Code:    constant.INVALID_PARAMS,
		Message: "Invalid params",
		Data:    nil,
	}
}

func ErrorResponse(message string) Respond {
	return Respond{
		Code:    constant.ERROR_RESPONSE,
		Message: message,
		Data:    nil,
	}
}

func InvalidCredential() Respond {
	return Respond{
		Code:    constant.INVALID_CREDENTIAL,
		Message: "Invalid credential",
		Data:    nil,
	}
}

func FieldAlreadyExist(entityName string) Respond {
	return Respond{
		Code:    constant.ALREADY_EXIST,
		Message: fmt.Sprintf("%v already existed", entityName),
		Data:    nil,
	}
}

func CanNotCreate(entityName string) Respond {
	return Respond{
		Code:    constant.ALREADY_EXIST,
		Message: fmt.Sprintf("Can not create %v", entityName),
		Data:    nil,
	}
}
