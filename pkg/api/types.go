package api

import "github.com/iwilltry42/norman/v3/pkg/types"

type ResponseWriter interface {
	Write(apiOp *types.APIRequest, code int, obj interface{})
}
