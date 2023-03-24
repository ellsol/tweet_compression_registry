package swagger

type ErrorResponse struct {
	// enum: Unknown,Unexpected,Unauthorized,Forbidden,Validation,BadRequest,NotFound,Conflict
	Kind          string               `json:"kind"`
	Code          string               `json:"code"` // a unique code matching [a-Z0-9]+
	Message       string               `json:"message"`
	Details       string               `json:"details"`
	InvalidFields map[string][2]string `json:"invalidFields,omitempty"` // field name -> [code, message]
}

// swagger:response BadRequestResponse
type BadRequestResponse struct {
	// in:body
	Body ErrorResponse
}

// swagger:response ForbiddenResponse
type ForbiddenResponse struct {
	// in:body
	Body ErrorResponse
}

// swagger:response NotFoundResponse
type NotFoundResponse struct {
	// in:body
	Body ErrorResponse
}

// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	// in:body
	Body ErrorResponse
}

// swagger:response UnauthorizedResponse
type UnauthorizedResponse struct {
	// in:body
	Body ErrorResponse
}

// swagger:response NoContent
type NoContent struct{}
