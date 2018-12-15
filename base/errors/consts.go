package errors

var (
	BadRequest            = Error{400, "BadRequest", "Bad Request"}
	InvalidParameter      = Error{400, "InvalidParameter", "A parameter specified in a request is not valid, is unsupported, or cannot be used."}
	InvalidQueryParameter = Error{400, "InvalidQueryParameter", "The query string is malformed or does not adhere to service standards."}
	MalformedParameter    = Error{400, "MalformedParameter", "The parameter specified in a request is not valid, is contains a syntax error, or cannot be decoded."}
	Unauthorized          = Error{401, "Unauthorized", "Unauthorized"}
	AccessDenied          = Error{403, "AccessDenied", "Access Denied"}
	InternalError         = Error{500, "InternalError", "An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the service forums."}
)
