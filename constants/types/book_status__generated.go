package types

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

var InvalidBookStatus = errors.New("invalid BookStatus")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("BookStatus", map[string]string{
		"NORMAL":   "正常展示",
		"PENGDING": "等待导入",
		"PROCESS":  "导入中",
		"READY":    "就绪",
	})
}

func ParseBookStatusFromString(s string) (BookStatus, error) {
	switch s {
	case "":
		return BOOK_STATUS_UNKNOWN, nil
	case "NORMAL":
		return BOOK_STATUS__NORMAL, nil
	case "PENGDING":
		return BOOK_STATUS__PENGDING, nil
	case "PROCESS":
		return BOOK_STATUS__PROCESS, nil
	case "READY":
		return BOOK_STATUS__READY, nil
	}
	return BOOK_STATUS_UNKNOWN, InvalidBookStatus
}

func ParseBookStatusFromLabelString(s string) (BookStatus, error) {
	switch s {
	case "":
		return BOOK_STATUS_UNKNOWN, nil
	case "正常展示":
		return BOOK_STATUS__NORMAL, nil
	case "等待导入":
		return BOOK_STATUS__PENGDING, nil
	case "导入中":
		return BOOK_STATUS__PROCESS, nil
	case "就绪":
		return BOOK_STATUS__READY, nil
	}
	return BOOK_STATUS_UNKNOWN, InvalidBookStatus
}

func (BookStatus) EnumType() string {
	return "BookStatus"
}

func (BookStatus) Enums() map[int][]string {
	return map[int][]string{
		int(BOOK_STATUS__NORMAL):   {"NORMAL", "正常展示"},
		int(BOOK_STATUS__PENGDING): {"PENGDING", "等待导入"},
		int(BOOK_STATUS__PROCESS):  {"PROCESS", "导入中"},
		int(BOOK_STATUS__READY):    {"READY", "就绪"},
	}
}
func (v BookStatus) String() string {
	switch v {
	case BOOK_STATUS_UNKNOWN:
		return ""
	case BOOK_STATUS__NORMAL:
		return "NORMAL"
	case BOOK_STATUS__PENGDING:
		return "PENGDING"
	case BOOK_STATUS__PROCESS:
		return "PROCESS"
	case BOOK_STATUS__READY:
		return "READY"
	}
	return "UNKNOWN"
}

func (v BookStatus) Label() string {
	switch v {
	case BOOK_STATUS_UNKNOWN:
		return ""
	case BOOK_STATUS__NORMAL:
		return "正常展示"
	case BOOK_STATUS__PENGDING:
		return "等待导入"
	case BOOK_STATUS__PROCESS:
		return "导入中"
	case BOOK_STATUS__READY:
		return "就绪"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BookStatus)(nil)

func (v BookStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBookStatus
	}
	return []byte(str), nil
}

func (v *BookStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBookStatusFromString(string(bytes.ToUpper(data)))
	return
}
