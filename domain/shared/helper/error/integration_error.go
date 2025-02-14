package error

import "fmt"

type ErrIntegration int

const (
	ERROR_NOT_AUTHORIZED ErrIntegration = iota
	ERROR_DATA_NOT_FOUND
	ERROR_API_INTEGRATION
	ERROR_NETWORK_CONNECTION
	ERROR_REQUEST_TIMEOUT
	ERROR_INTERNAL_SERVER
)

type IntegrationError struct {
	Type    ErrIntegration
	Message string
}

func (e IntegrationError) Error() string {
	return fmt.Sprintf("🔥 [Error] %s", e.Message)
}

func NewIntegrationError(errType ErrIntegration, message string) IntegrationError {
	return IntegrationError{
		Type:    errType,
		Message: message,
	}
}
