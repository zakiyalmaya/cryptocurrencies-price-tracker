package exchangerate

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=ExchangeRateService.go
type ExchangeRateService interface {
	GetLatest(base, to string) (float64, error)
}