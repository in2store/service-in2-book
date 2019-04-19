package types

//go:generate libtools gen enum BookStatus
// swagger:enum
type BookStatus uint8

// 书籍状态
const (
	BOOK_STATUS_UNKNOWN   BookStatus = iota
	BOOK_STATUS__PENGDING            // 等待导入
	BOOK_STATUS__PROCESS             // 导入中
	BOOK_STATUS__READY               // 就绪
	BOOK_STATUS__NORMAL              // 正常展示
)
