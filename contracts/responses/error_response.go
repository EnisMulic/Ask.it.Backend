package responses

type ErrorResponseModel struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

type ErrorResponse struct {
	Errors []ErrorResponseModel `json:"errors"`
}

func NewErrorResponse(err ErrorResponseModel) *ErrorResponse {
	var errors []ErrorResponseModel
	errors = append(errors, err)

	return &ErrorResponse{errors}
}

func (errors *ErrorResponse) AddError(err ErrorResponseModel) {
	errors.Errors = append(errors.Errors, err)
}