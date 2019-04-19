package types

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

var InvalidBookLanguage = errors.New("invalid BookLanguage")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("BookLanguage", map[string]string{
		"ENGLISH": "英文",
		"ZH_CN":   "简体中文",
		"ZH_TW":   "繁体中文",
	})
}

func ParseBookLanguageFromString(s string) (BookLanguage, error) {
	switch s {
	case "":
		return BOOK_LANGUAGE_UNKNOWN, nil
	case "ENGLISH":
		return BOOK_LANGUAGE__ENGLISH, nil
	case "ZH_CN":
		return BOOK_LANGUAGE__ZH_CN, nil
	case "ZH_TW":
		return BOOK_LANGUAGE__ZH_TW, nil
	}
	return BOOK_LANGUAGE_UNKNOWN, InvalidBookLanguage
}

func ParseBookLanguageFromLabelString(s string) (BookLanguage, error) {
	switch s {
	case "":
		return BOOK_LANGUAGE_UNKNOWN, nil
	case "英文":
		return BOOK_LANGUAGE__ENGLISH, nil
	case "简体中文":
		return BOOK_LANGUAGE__ZH_CN, nil
	case "繁体中文":
		return BOOK_LANGUAGE__ZH_TW, nil
	}
	return BOOK_LANGUAGE_UNKNOWN, InvalidBookLanguage
}

func (BookLanguage) EnumType() string {
	return "BookLanguage"
}

func (BookLanguage) Enums() map[int][]string {
	return map[int][]string{
		int(BOOK_LANGUAGE__ENGLISH): {"ENGLISH", "英文"},
		int(BOOK_LANGUAGE__ZH_CN):   {"ZH_CN", "简体中文"},
		int(BOOK_LANGUAGE__ZH_TW):   {"ZH_TW", "繁体中文"},
	}
}
func (v BookLanguage) String() string {
	switch v {
	case BOOK_LANGUAGE_UNKNOWN:
		return ""
	case BOOK_LANGUAGE__ENGLISH:
		return "ENGLISH"
	case BOOK_LANGUAGE__ZH_CN:
		return "ZH_CN"
	case BOOK_LANGUAGE__ZH_TW:
		return "ZH_TW"
	}
	return "UNKNOWN"
}

func (v BookLanguage) Label() string {
	switch v {
	case BOOK_LANGUAGE_UNKNOWN:
		return ""
	case BOOK_LANGUAGE__ENGLISH:
		return "英文"
	case BOOK_LANGUAGE__ZH_CN:
		return "简体中文"
	case BOOK_LANGUAGE__ZH_TW:
		return "繁体中文"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BookLanguage)(nil)

func (v BookLanguage) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBookLanguage
	}
	return []byte(str), nil
}

func (v *BookLanguage) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBookLanguageFromString(string(bytes.ToUpper(data)))
	return
}
