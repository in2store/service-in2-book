package types

//go:generate libtools gen enum CodeLanguage
// swagger:enum
type CodeLanguage uint8

// 代码语言
const (
	CODE_LANGUAGE_UNKNOWN     CodeLanguage = iota
	CODE_LANGUAGE__C_CPP                   // C/C++
	CODE_LANGUAGE__JAVA                    // Java
	CODE_LANGUAGE__JAVASCRIPT              // Javascript
	CODE_LANGUAGE__PYTHON                  // Python
	CODE_LANGUAGE__CSHARP                  // C#
	CODE_LANGUAGE__GOLANG                  // Golang
	CODE_LANGUAGE__PHP                     // PHP
)
