package errorx

import "encoding/json"

type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ParseError(str string) *CodeErrorResponseContent {
	res := &CustomError{}
	if err := json.Unmarshal([]byte(str), res); err == nil {
		return &CodeErrorResponseContent{
			Code: ErrCodeInternal,
			Msg:  res.Message,
		}
	}
	return ErrCodeInternal.GenError(str)
}
