package service

// DifService
type DifService struct {
	difMessage string
}

// NewDifService creates new dif service instance
func NewDifService(difMessage string) *DifService {
	return &DifService{difMessage: difMessage}
}

// Initialize dif service
func (difService *DifService) Init(err error) {
	return
}
