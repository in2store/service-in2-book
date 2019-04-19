package types

//go:generate libtools gen enum BookLanguage
// swagger:enum
type BookLanguage uint8

// 文档语言
const (
	BOOK_LANGUAGE_UNKNOWN  BookLanguage = iota
	BOOK_LANGUAGE__ZH_CN                // 简体中文
	BOOK_LANGUAGE__ZH_TW                // 繁体中文
	BOOK_LANGUAGE__ENGLISH              // 英文
)
