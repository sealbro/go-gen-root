package services

type ServiceOne struct {
}

// NewServiceOne creates a new instance of ServiceOne
// @inject
func NewServiceOne() *ServiceOne {
	return &ServiceOne{}
}

func (s *ServiceOne) CallOne() int {
	return 1
}
