RESULT_PROJECT_NAME_EXISTpackage models

func GetErrorString(code int) string {
	switch code {
		case RESULT_OK:
			return "OK"
		case RESULT_ERROR:
			return "Error"
		case RESULT_NOT_LOGIN:
			return "Not logged in"
		case RESULT_TOKEN_ERROR:
			return "Token error"
		case RESULT_PASSWORD_ERROR:
			return "Password error"
		case RESULT_NO_PERMISSION:
			return "No Permission"
		case RESULT_PARAM_ERROR:
			return "Parameter error"
		case RESULT_USERNAME_EXIST:
			return "User name exists"
		case RESULT_PHONE_EXIST:
			return "Phone number exists"
		case RESULT_USER_NOT_EXIST:
			return "User does not exist"
		case RESULT_PROJECT_NOT_EXIST:
			return "Project does not exist"
		case RESULT_TOKEN_EXPIRED:
			return "Token expired"
		case RESULT_AUTH_ERROR:
			return "Auth error"
		case RESULT_RECORD_NOT_EXIST:
			return "Record not found"
		case RESULT_RECORD_EXIST:
			return  "Record exists"
		default:
			return "Unknown error"
	}
}

const (
	RESULT_OK = iota
	RESULT_ERROR
	RESULT_NOT_LOGIN
	RESULT_TOKEN_ERROR
	RESULT_PASSWORD_ERROR
	RESULT_NO_PERMISSION
	RESULT_USER_NOT_EXIST
	RESULT_PARAM_ERROR
	RESULT_USERNAME_EXIST
	RESULT_PHONE_EXIST
	RESULT_PROJECT_NOT_EXIST
	RESULT_TOKEN_EXPIRED
	RESULT_AUTH_ERROR
	RESULT_RECORD_NOT_EXIST
	RESULT_RECORD_EXIST
)

