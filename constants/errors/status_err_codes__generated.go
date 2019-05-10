package errors

import (
	"github.com/johnnyeven/libtools/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("BadRequest", 400099000, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("Unauthorized", 401099000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403099000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("NotFound", 404099000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("CategoryKeyNotFound", 404099001, "分类标识未找到", "", true)
	status_error.StatusErrorCodes.Register("BookNotFound", 404099002, "文档未找到", "", true)
	status_error.StatusErrorCodes.Register("TagNotFound", 404099003, "标签未找到", "", true)
	status_error.StatusErrorCodes.Register("Conflict", 409099000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("CategoryKeyConflict", 409099001, "分类标识已存在", "", true)
	status_error.StatusErrorCodes.Register("BookTagConflict", 409099002, "文档已在分类中", "", true)
	status_error.StatusErrorCodes.Register("BookConflict", 409099003, "文档已存在", "", true)
	status_error.StatusErrorCodes.Register("TagConflict", 409099004, "标签名称已存在", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500099000, "内部处理错误", "", false)
}
