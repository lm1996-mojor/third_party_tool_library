package third_party_tool_library

type ResponseResult struct {
	Code    *string
	Message *string
}

func NewResult(code *string, message *string) ResponseResult {
	return ResponseResult{code, message}
}
