package service

// AttService
type AttService struct {
	attMessage string
}

// NewAttService creates new att service instance
func NewAttService(attMessage string) *AttService {
	return &AttService{attMessage: attMessage}
}

// Initialize att service
func (attService *AttService) Init(err error) {
	return
}
