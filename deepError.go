package deepError

import "fmt"

// BuildVersion is the build version of the entire project
var BuildVersion string = "2.0"

// DeepErr is the actual error struct to be used
type DeepErr struct {
	Function string
	Action   string
	Code     string
	Message  string
	Cause    error
	Params   []interface{}
}

// New is a basic function which gives info on the error. In Which function, doing what action, and what was the nested error.
func New(function, action string, cause error) DeepErr {
	return DeepErr{
		Function: function,
		Action:   action,
		Cause:    cause,
	}
}

// NewFull is the full error function with all params
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
