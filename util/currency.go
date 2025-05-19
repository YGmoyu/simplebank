package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	RMB = "RMB"
)

// 检查输入的货币是否支持
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, RMB:
		return true
	}
	return false
}
