package businesslogic

import "github.com/sealbro/go-gen-root/examples/app-with-deps/internal/services"

type CalculatorService struct {
	one *services.ServiceOne
	two *services.ServiceTwo
}

// NewCalculatorService creates a new instance of CalculatorService
// @inject
func NewCalculatorService(one *services.ServiceOne, two *services.ServiceTwo) *CalculatorService {
	return &CalculatorService{
		one: one,
		two: two,
	}
}

func (s *CalculatorService) Calculate() int {
	return s.one.CallOne() + s.two.CallTwo()
}
