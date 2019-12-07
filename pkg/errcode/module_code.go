package errcode

var (
	ERROR_GET_TAG_LIST_FAIL = NewError(20010001, "获取标签列表失败")
	ERROR_CREATE_TAG_FAIL   = NewError(20010002, "创建标签失败")
	ERROR_UPDATE_TAG_FAIL   = NewError(20010003, "更新标签失败")
	ERROR_DELETE_TAG_FAIL   = NewError(20010004, "删除标签失败")
	ERROR_COUNT_TAG_FAIL    = NewError(20010005, "统计标签失败")

	ERROR_GET_ARTICLE_FAIL    = NewError(20020001, "获取单个文章失败")
	ERROR_GET_ARTICLES_FAIL   = NewError(20020002, "获取多个文章失败")
	ERROR_CREATE_ARTICLE_FAIL = NewError(20020003, "创建文章失败")
	ERROR_UPDATE_ARTICLE_FAIL = NewError(20020004, "更新文章失败")
	ERROR_DELETE_ARTICLE_FAIL = NewError(20020005, "删除文章失败")

	ERROR_UPLOAD_FILE_FAIL = NewError(20030001, "上传文件失败")
)
