package exchangerate

type ExchangeRateService interface {
	GetLatest(base, to string) (float64, error)
}