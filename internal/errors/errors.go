package errors

import (
	"errors"

	"git.internal.yunify.com/manage/common/gerr"
)

/*

	关于错误的层次设计：
		从grpc service返回错误。


*/

// func NewErrorCode(code int32, msg string) error {
// 	return formatCodeError(code, msg)
// }

// func NewError(msg string) error {
// 	return status.Error(codes.OK, msg)
// }

// func formatCodeError(code int32, msg string) error {
// 	return status.Error(codes.Code(code), msg)
// }

type Error struct {
	code    string
	message string
}

func New(code string, err error) *Error {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return NewError(code, msg)
}

func NewError(code, msg string) *Error {
	return &Error{
		code:    code,
		message: msg,
	}
}

func (this *Error) Error() string {
	if nil == this {
		return "Nil"
	}
	val := this.message
	if CodeEmpty != this.code {
		val = this.code
	}
	return val
}

func (this *Error) Err() error {
	if this == nil {
		return nil
	}
	val := this.message
	if "" != this.code {
		val = this.code
	}
	return errors.New(val)
}

func ConvertError(err error) error {
	if nil == err {
		return err
	}
	switch err.(type) {
	case *Error:
		er, _ := err.(*Error)
		err = er.Err()
	default:
	}
	return err
}

func (this *Error) Code() string {
	if nil == this {
		return ""
	}
	return this.code
}

func (this *Error) Message() string {
	if nil == this {
		return ""
	}
	return this.message
}

//-----------------------------------------Errors List---------------------------------------
var (
	CodeEmpty     = ""
	InternalError = gerr.InternalError

	RuleNotExisted      = gerr.RuleNotExisted
	RuleStatusInvalid   = gerr.RuleStatusInvalid
	RuleOperatorInvalid = gerr.RuleOperatorInvalid
	RuleRequierdId      = gerr.RuleRequierdId
	RuleLimit           = gerr.RuleLimit
	RuleConfigInvalid   = gerr.RuleConfigInvalid
	RuleBadArgument     = gerr.RuleBadArgument
	RuleParseError      = gerr.RuleParseError
	RuleValidateFailed  = gerr.RuleValidateFailed

	//Action
	RuleActionNotExisted    = gerr.RuleActionNotExisted
	RuleActionTypeInvalid   = gerr.RuleActionTypeInvalid
	RuleActionNotBindTable  = gerr.RuleActionNotBindTable
	RuleActionConfigInvalid = gerr.RuleActionConfigInvalid
	RuleActionNonAvailable  = gerr.RuleActionNonAvailable
	RuleActionLimit         = gerr.RuleActionLimit

	//Sink                   = Action + "Sink."
	SinkMissingParams         = gerr.SinkMissingParams
	SinkEndpointsEmpty        = gerr.SinkEndpointsEmpty
	SinkTypeNotSupport        = gerr.SinkTypeNotSupport
	SinkTableNotExisted       = gerr.SinkTableNotExisted
	SinkConnectError          = gerr.SinkConnectError
	SinkUnkownTableField      = gerr.SinkUnkownTableField
	SinkTableInvalidFieldType = gerr.SinkTableInvalidFieldType
	SinkNotVerify             = gerr.SinkNotVerify
	SinkMappingNotExisted     = gerr.SinkMappingNotExisted

	RuleStatusMessage = gerr.RuleStatusMessage
)
