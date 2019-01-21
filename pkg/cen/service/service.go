package service

// CenService
type CenService struct {
	cenMessage string
}

// NewCenService creates new cen service instance
func NewCenService(cenMessage string) *CenService {
	return &CenService{cenMessage: cenMessage}
}

// Initialize cen service
func (cenService *CenService) Init(err error) {
	return
}
