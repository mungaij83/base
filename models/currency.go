package models

import "time"

type CurrencyModel struct {
	Name          string  //: fields.Char{String: "Currency", Help: "Currency Code [ISO 4217]", Size: 3, Unique: true},
	Symbol        string  // ": fields.Char{Help: "Currency sign, to be used when printing amounts", Size: 4},
	Rate          float32 // "Rate": fields.Float{String: "Current Rate",  Help: "The rate of the currency to the currency of rate 1", Digits: nbutils.Digits{Precision: 16, Scale: 6}, Compute float64 //: h.Currency().Methods().ComputeCurrentRate(), Depends: []string{"Rates", "Rates.Rate"}},
	Rates         float64 //": fields.One2Many{RelationModel: h.CurrencyRate(), ReverseFK: "Currency"},
	Rounding      float32 //fields.Float{String: "Rounding Factor", Digits: nbutils.Digits{Precision: 12, Scale: 6}},
	DecimalPlaces int     //": fields.Integer{GoType: new(int), Compute: h.Currency().Methods().ComputeDecimalPlaces(), Depends: []string{"Rounding"}},
	Active        bool
	Position      string    //: fields.Selection{Selection: types.Selection{"after": "After Amount", "before": "Before Amount"}, String: "Symbol Position", Help: "Determines where the currency symbol should be placed after or before the amount."},
	Date          time.Time //": fields.Date{Compute: h.Currency().Methods().ComputeDate(), Depends: []string{"Rates", "Rates.Name"}},
}

func (CurrencyModel) TableName() string {
	return "Currency"
}


