package errors

import (
	"net/http"

	"github.com/johnnyeven/libtools/courier/status_error"
)

//go:generate libtools gen error
const ServiceStatusErrorCode = 99 * 1e3 // todo rename this

const (
	// 请求参数错误
	BadRequest status_error.StatusErrorCode = http.StatusBadRequest*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 未找到
	NotFound status_error.StatusErrorCode = http.StatusNotFound*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 分类标识未找到
	CategoryKeyNotFound
	// @errTalk 文档未找到
	BookNotFound
	// @errTalk 标签未找到
	TagNotFound
)

const (
	// @errTalk 未授权
	Unauthorized status_error.StatusErrorCode = http.StatusUnauthorized*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 操作冲突
	Conflict status_error.StatusErrorCode = http.StatusConflict*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 分类标识已存在
	CategoryKeyConflict
	// @errTalk 文档已在分类中
	BookTagConflict
	// @errTalk 文档已存在
	BookConflict
	// @errTalk 标签名称已存在
	TagConflict
)

const (
	// @errTalk 不允许操作
	Forbidden status_error.StatusErrorCode = http.StatusForbidden*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 内部处理错误
	InternalError status_error.StatusErrorCode = http.StatusInternalServerError*1e6 + ServiceStatusErrorCode + iota
)
