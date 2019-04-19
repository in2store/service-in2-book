package types

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

var InvalidCodeLanguage = errors.New("invalid CodeLanguage")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("CodeLanguage", map[string]string{
		"CSHARP":     "C#",
		"C_CPP":      "C/C++",
		"GOLANG":     "Golang",
		"JAVA":       "Java",
		"JAVASCRIPT": "Javascript",
		"PHP":        "PHP",
		"PYTHON":     "Python",
	})
}

func ParseCodeLanguageFromString(s string) (CodeLanguage, error) {
	switch s {
	case "":
		return CODE_LANGUAGE_UNKNOWN, nil
	case "CSHARP":
		return CODE_LANGUAGE__CSHARP, nil
	case "C_CPP":
		return CODE_LANGUAGE__C_CPP, nil
	case "GOLANG":
		return CODE_LANGUAGE__GOLANG, nil
	case "JAVA":
		return CODE_LANGUAGE__JAVA, nil
	case "JAVASCRIPT":
		return CODE_LANGUAGE__JAVASCRIPT, nil
	case "PHP":
		return CODE_LANGUAGE__PHP, nil
	case "PYTHON":
		return CODE_LANGUAGE__PYTHON, nil
	}
	return CODE_LANGUAGE_UNKNOWN, InvalidCodeLanguage
}

func ParseCodeLanguageFromLabelString(s string) (CodeLanguage, error) {
	switch s {
	case "":
		return CODE_LANGUAGE_UNKNOWN, nil
	case "C#":
		return CODE_LANGUAGE__CSHARP, nil
	case "C/C++":
		return CODE_LANGUAGE__C_CPP, nil
	case "Golang":
		return CODE_LANGUAGE__GOLANG, nil
	case "Java":
		return CODE_LANGUAGE__JAVA, nil
	case "Javascript":
		return CODE_LANGUAGE__JAVASCRIPT, nil
	case "PHP":
		return CODE_LANGUAGE__PHP, nil
	case "Python":
		return CODE_LANGUAGE__PYTHON, nil
	}
	return CODE_LANGUAGE_UNKNOWN, InvalidCodeLanguage
}

func (CodeLanguage) EnumType() string {
	return "CodeLanguage"
}

func (CodeLanguage) Enums() map[int][]string {
	return map[int][]string{
		int(CODE_LANGUAGE__CSHARP):     {"CSHARP", "C#"},
		int(CODE_LANGUAGE__C_CPP):      {"C_CPP", "C/C++"},
		int(CODE_LANGUAGE__GOLANG):     {"GOLANG", "Golang"},
		int(CODE_LANGUAGE__JAVA):       {"JAVA", "Java"},
		int(CODE_LANGUAGE__JAVASCRIPT): {"JAVASCRIPT", "Javascript"},
		int(CODE_LANGUAGE__PHP):        {"PHP", "PHP"},
		int(CODE_LANGUAGE__PYTHON):     {"PYTHON", "Python"},
	}
}
func (v CodeLanguage) String() string {
	switch v {
	case CODE_LANGUAGE_UNKNOWN:
		return ""
	case CODE_LANGUAGE__CSHARP:
		return "CSHARP"
	case CODE_LANGUAGE__C_CPP:
		return "C_CPP"
	case CODE_LANGUAGE__GOLANG:
		return "GOLANG"
	case CODE_LANGUAGE__JAVA:
		return "JAVA"
	case CODE_LANGUAGE__JAVASCRIPT:
		return "JAVASCRIPT"
	case CODE_LANGUAGE__PHP:
		return "PHP"
	case CODE_LANGUAGE__PYTHON:
		return "PYTHON"
	}
	return "UNKNOWN"
}

func (v CodeLanguage) Label() string {
	switch v {
	case CODE_LANGUAGE_UNKNOWN:
		return ""
	case CODE_LANGUAGE__CSHARP:
		return "C#"
	case CODE_LANGUAGE__C_CPP:
		return "C/C++"
	case CODE_LANGUAGE__GOLANG:
		return "Golang"
	case CODE_LANGUAGE__JAVA:
		return "Java"
	case CODE_LANGUAGE__JAVASCRIPT:
		return "Javascript"
	case CODE_LANGUAGE__PHP:
		return "PHP"
	case CODE_LANGUAGE__PYTHON:
		return "Python"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CodeLanguage)(nil)

func (v CodeLanguage) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCodeLanguage
	}
	return []byte(str), nil
}

func (v *CodeLanguage) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCodeLanguageFromString(string(bytes.ToUpper(data)))
	return
}
