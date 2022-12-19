package models

type CurrencyRateModel struct {
	Name     string        `hexya:"required;index"`
	Rate     float64       `hexya:"precision=16;scale=6;help=The rate of the currency to the currency of rate 1"`
	Currency CurrencyModel `hexya:"many2one=id"`
	Company  CompanyModel  `hexya:"many2one=id"`
}
