package utils

import "github.com/EnisMulic/Ask.it.Backend/contracts/responses"

func GetIntPointer(value int64) *int64 {
    return &value
}

func CreateQuestionPagedResponse(data []responses.QuestionResponseModel, count int64, pageNumber int64, pageSize int64) *responses.QuestionsReponse{
	response := responses.QuestionsReponse{
		Data: data,
	}

	if pageSize > 1 {
		response.LastPage = count / pageSize
	} else {
		response.LastPage = 1
	}

	if pageNumber >= 1 && pageNumber < response.LastPage && len(data) > 0 {
		response.NextPage = GetIntPointer(pageNumber + 1)
	} else {
		response.NextPage = nil
	}

	if pageNumber - 1 >= 1 {
		response.PreviousPage = GetIntPointer(pageNumber - 1)
	} else {
		response.PreviousPage = nil
	}

	if pageNumber >= 1 {
		response.PageNumber = pageNumber
	} else {
		response.PageNumber = 1
	}

	if pageSize >= 1 {
		response.PageSize = pageSize
	} else {
		response.PageSize = count
	}

	response.FirstPage = 1

	return &response
}

func CreateUserPagedResponse(data []responses.UserResponseModel, count int64, pageNumber int64, pageSize int64) responses.UsersResponse{
	response := responses.UsersResponse {
		Data: data,
	}
	
	if pageSize > 1 {
		response.LastPage = count / pageSize
	} else {
		response.LastPage = 1
	}

	if pageNumber > 1 && pageNumber < response.LastPage && len(data) > 0 {
		response.NextPage = GetIntPointer(pageNumber + 1)
	} else {
		response.NextPage = nil
	}

	if pageNumber - 1 >= 1 {
		response.PreviousPage = GetIntPointer(pageNumber - 1)
	} else {
		response.PreviousPage = nil
	}

	if pageNumber >= 1 {
		response.PageNumber = pageNumber
	} else {
		response.PageNumber = 1
	}

	if pageSize >= 1 {
		response.PageSize = pageSize
	} else {
		response.PageSize = count
	}

	response.FirstPage = 1

	return response
}