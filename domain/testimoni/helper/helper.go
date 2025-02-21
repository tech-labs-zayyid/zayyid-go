package helper

import (
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/testimoni/model"
)

type nilByteRequest struct {
}

func CheckRequest(request model.User) (response model.User, err error) {
	if request.Id == "" {
		return request, sharedError.SetErrorMessage("email")
	}

	if request.Name == "" {
		return request, sharedError.SetErrorMessage("name")
	}

	response = request
	return
}
