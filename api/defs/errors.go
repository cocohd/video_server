package defs

// 自定义error数据结构

type Err struct {
	Error string `json:"error"`
	// error编码
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	// 请求中无内容
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:     "Request body is not correct",
			ErrorCode: "001",
		},
	}

	// 无权限
	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error:     "User has no authentication",
			ErrorCode: "002",
		},
	}

	// 定义数据库写入错误
	ErrorDBError = ErrorResponse{
		HttpSC: 501,
		Error: Err{
			Error:     "DB ops failed",
			ErrorCode: "003",
		},
	}

	// 定义internal错误
	ErrorInternalError = ErrorResponse{
		HttpSC: 501,
		Error: Err{
			Error:     "internal error",
			ErrorCode: "004",
		},
	}
)
