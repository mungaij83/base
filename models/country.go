package models

type CountryGroupModel struct {
	Name      string         `hexya:"required"`
	Countries []CountryModel `hexya:"many2many"`
}

type CountryModel struct {
	Name          string // fields.Char{String: "Country Name", Help: "The full name of the country.", Translate: true, Required: true, Unique: true},
	Code          string //fields.Char{String: "Country Code", Size: 2, Unique: true, Help: "The ISO country code in two chars.\nYou can use this field for quick search."},
	AddressFormat string
	AddressViewID string
	Currency      CurrencyModel //   fields.Many2One{RelationModel: h.Currency()},
	Image         []byte
	PhoneCode     int                 //fields.Integer{String: "Country Calling Code"},
	CountryGroups CountryGroupModel   //: fields.Many2Many{RelationModel: h.CountryGroup()},
	States        []CountryStateModel //fields.One2Many{RelationModel: h.CountryState(), ReverseFK: "Country"},
	NamePosition  string
	VATLabel      string //": fields.Char{Translate: true, Help: "Use this field if you want to change vat label."},
}

type CountryStateModel struct {
	Name    string       `hexya:"string=State Name;required;help=Administrative divisions of a country. E.g. Fed. State, Departement, Canton"` //String: "State Name", Required: true,Administrative divisions of a country. E.g. Fed. State, Departement, Canton
	Code    string       `hexya:"string=State Code;size=3;help=The state code in max. three chars.;required"`                                  //{String: "State Code", Size: 3, Help: "The state code in max. three chars.", Required: true},
	Country CountryModel `hexya:"required"`                                                                                                    //{RelationModel: h.Country(), Required: true},
}
