package models
type AnswerMessage struct {
	Status  bool `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int `json:"code"`
}

func GetAnswerByErrorCode(errcode int) AnswerMessage {
	var answermessage AnswerMessage
	answermessage.Message = GetErrorString(errcode)
	answermessage.Status = false
	answermessage.Code = errcode
	return answermessage
}

func GetSuccessAnswerWithParam(param interface{}) AnswerMessage {
	var answermessage AnswerMessage
	answermessage.Data = param
	answermessage.Status = true
	return answermessage
}

func GetAnswer(errcode int, param interface{}) interface{} {
	if errcode == RESULT_OK {
		if param == nil {
			return GetSuccessAnswer()
		} else {
			return GetSuccessAnswerWithParam(param)
		}
	} else {
		return GetAnswerByErrorCode(errcode)
	}
}

