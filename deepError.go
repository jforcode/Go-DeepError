package deepError

import "fmt"

type DeepErr struct {
	Function string
	Action   string
	Code     string
	Message  string
	Cause    error
	Params   []interface{}
}

func New(function, action string, cause error) DeepErr {
	return DeepErr{
		Function: function,
		Action:   action,
		Cause:    cause,
	}
}

func NewFull(function, action string, cause error, code, message string, params []interface{}) DeepErr {
	return DeepErr{
		Function: function,
		Action:   action,
		Cause:    cause,
		Code:     code,
		Message:  message,
	}
}

func (err DeepErr) Error() string {
	ret := ""
	if err.Function != "" {
		ret += "IN " + err.Function
	} else {
		ret += "IN <NO_FUNCTION>"
	}

	if err.Action != "" {
		ret += " WHILE " + err.Action
	} else {
		ret += " WHILE <NO_ACTION>"
	}

	if err.Code != "" {
		ret += " GOT " + err.Code
	} else {
		ret += " GOT <NO_CODE>"
	}

	if err.Message != "" {
		ret += " (" + err.Message + ")"
	} else {
		ret += " (<NO_MESSAGE>)"
	}

	if err.Params != nil && len(err.Params) != 0 {
		ret += ": " + fmt.Sprint(err.Params)
	} else {
		ret += ": <NO_PARAMS>"
	}

	if err.Cause != nil {
		ret += "\nNested Error: " + err.Cause.Error()
	}

	return ret
}
