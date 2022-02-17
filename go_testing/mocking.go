package go_testing

import "fmt"

type SumClient interface {
	Sum(a, b int) (int, error)
}

type SumService struct {
	sumClient SumClient
}

func NewSumService(sumClient SumClient) SumService {
	return SumService{
		sumClient: sumClient,
	}
}

func (ss SumService) SumWithExternalCall(a, b int) (int, error) {
	result, err := ss.sumClient.Sum(a, b)
	if err != nil {
		return 0, fmt.Errorf("something happens calling external service: %w", err)
	}

	return result, nil
}
