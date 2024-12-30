package services

type ServiceTwo struct {
}

// NewServiceTwo creates a new instance of ServiceTwo
// @inject
func NewServiceTwo() *ServiceTwo {
	return &ServiceTwo{}
}

func (s *ServiceTwo) CallTwo() int {
	return 2
}
