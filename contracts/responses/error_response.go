package responses

type ErrorResponseModel struct {
	FieldName string
	Message   string
}

type ErrorResponse struct {
	Errors []ErrorResponseModel
}

func NewErrorResponse(err ErrorResponseModel) *ErrorResponse {
	var errors []ErrorResponseModel
	errors = append(errors, err)

	return &ErrorResponse{errors}
}

func (errors *ErrorResponse) AddError(err ErrorResponseModel) {
	errors.Errors = append(errors.Errors, err)
}