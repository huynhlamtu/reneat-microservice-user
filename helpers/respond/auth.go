package respond

import "reneat-microservice-user/constant"

func MissingRegister() Respond {
	return Respond{
		Code:    constant.MISSING_REGISTER_INFO,
		Message: "Register info is missing",
		Data:    nil,
	}
}

func MissingLogin() Respond {
	return Respond{
		Code:    constant.MISSING_USERNAME_PASSWORD,
		Message: "Username or password is missing",
		Data:    nil,
	}
}

func EmailPasswordIncorrect() Respond {
	return Respond{
		Code:    constant.USERNAME_PASSWORD_INCORRECT,
		Message: "Username and password incorrect",
		Data:    nil,
	}
}
