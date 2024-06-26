// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat.proto

package chat_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateIn with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateIn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateIn with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateInMultiError, or nil
// if none found.
func (m *CreateIn) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateIn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetUserIds()) < 1 {
		err := CreateInValidationError{
			field:  "UserIds",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_CreateIn_UserIds_Unique := make(map[int64]struct{}, len(m.GetUserIds()))

	for idx, item := range m.GetUserIds() {
		_, _ = idx, item

		if _, exists := _CreateIn_UserIds_Unique[item]; exists {
			err := CreateInValidationError{
				field:  fmt.Sprintf("UserIds[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_CreateIn_UserIds_Unique[item] = struct{}{}
		}

		if item <= 0 {
			err := CreateInValidationError{
				field:  fmt.Sprintf("UserIds[%v]", idx),
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CreateInMultiError(errors)
	}

	return nil
}

// CreateInMultiError is an error wrapping multiple validation errors returned
// by CreateIn.ValidateAll() if the designated constraints aren't met.
type CreateInMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateInMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateInMultiError) AllErrors() []error { return m }

// CreateInValidationError is the validation error returned by
// CreateIn.Validate if the designated constraints aren't met.
type CreateInValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateInValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateInValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateInValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateInValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateInValidationError) ErrorName() string { return "CreateInValidationError" }

// Error satisfies the builtin error interface
func (e CreateInValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateInValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateInValidationError{}

// Validate checks the field values on CreateOut with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateOut) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOut with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateOutMultiError, or nil
// if none found.
func (m *CreateOut) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOut) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateOutMultiError(errors)
	}

	return nil
}

// CreateOutMultiError is an error wrapping multiple validation errors returned
// by CreateOut.ValidateAll() if the designated constraints aren't met.
type CreateOutMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOutMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOutMultiError) AllErrors() []error { return m }

// CreateOutValidationError is the validation error returned by
// CreateOut.Validate if the designated constraints aren't met.
type CreateOutValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOutValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOutValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOutValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOutValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOutValidationError) ErrorName() string { return "CreateOutValidationError" }

// Error satisfies the builtin error interface
func (e CreateOutValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOut.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOutValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOutValidationError{}

// Validate checks the field values on DeleteIn with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteIn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteIn with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteInMultiError, or nil
// if none found.
func (m *DeleteIn) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteIn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteInValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteInMultiError(errors)
	}

	return nil
}

// DeleteInMultiError is an error wrapping multiple validation errors returned
// by DeleteIn.ValidateAll() if the designated constraints aren't met.
type DeleteInMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteInMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteInMultiError) AllErrors() []error { return m }

// DeleteInValidationError is the validation error returned by
// DeleteIn.Validate if the designated constraints aren't met.
type DeleteInValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteInValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteInValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteInValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteInValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteInValidationError) ErrorName() string { return "DeleteInValidationError" }

// Error satisfies the builtin error interface
func (e DeleteInValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteIn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteInValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteInValidationError{}

// Validate checks the field values on SendMessageIn with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendMessageIn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageIn with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendMessageInMultiError, or
// nil if none found.
func (m *SendMessageIn) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageIn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	// no validation rules for Text

	if len(errors) > 0 {
		return SendMessageInMultiError(errors)
	}

	return nil
}

// SendMessageInMultiError is an error wrapping multiple validation errors
// returned by SendMessageIn.ValidateAll() if the designated constraints
// aren't met.
type SendMessageInMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageInMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageInMultiError) AllErrors() []error { return m }

// SendMessageInValidationError is the validation error returned by
// SendMessageIn.Validate if the designated constraints aren't met.
type SendMessageInValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageInValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageInValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageInValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageInValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageInValidationError) ErrorName() string { return "SendMessageInValidationError" }

// Error satisfies the builtin error interface
func (e SendMessageInValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageIn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageInValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageInValidationError{}

// Validate checks the field values on ConnectChatIn with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConnectChatIn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectChatIn with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConnectChatInMultiError, or
// nil if none found.
func (m *ConnectChatIn) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectChatIn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	if len(errors) > 0 {
		return ConnectChatInMultiError(errors)
	}

	return nil
}

// ConnectChatInMultiError is an error wrapping multiple validation errors
// returned by ConnectChatIn.ValidateAll() if the designated constraints
// aren't met.
type ConnectChatInMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectChatInMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectChatInMultiError) AllErrors() []error { return m }

// ConnectChatInValidationError is the validation error returned by
// ConnectChatIn.Validate if the designated constraints aren't met.
type ConnectChatInValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectChatInValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectChatInValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectChatInValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectChatInValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectChatInValidationError) ErrorName() string { return "ConnectChatInValidationError" }

// Error satisfies the builtin error interface
func (e ConnectChatInValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectChatIn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectChatInValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectChatInValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for Text

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}

	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}
