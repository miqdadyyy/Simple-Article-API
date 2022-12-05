package constant

const (
	HTTPStatusOK                  = 200
	HTTPStatusBadRequest          = 400
	HTTPStatusUnauthorized        = 401
	HTTPStatusForbidden           = 403
	HTTPStatusNotFound            = 404
	HTTPStatusInternalServerError = 500
)

var (
	HTTPStatusMessageMap = map[int]string{
		HTTPStatusOK:                  "Success",
		HTTPStatusBadRequest:          "Bad Request",
		HTTPStatusUnauthorized:        "Unauthorized",
		HTTPStatusForbidden:           "Forbidden",
		HTTPStatusNotFound:            "Not Found",
		HTTPStatusInternalServerError: "Internal Server Error",
	}
)
