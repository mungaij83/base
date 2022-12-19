package models

type BankAccountModel struct {
	AccountType            string         //: fields.Char{Compute: h.BankAccount().Methods().ComputeAccountType(), Depends: []string{""}},
	Name                   string         `hexya:"display_name=Account Number;required"`
	SanitizedAccountNumber string         `hexya:"display_name=Account Number;required;stored"`
	Partner                *PartnerModel  `hexya:"display_name=Account Holder;required;index"`
	Bank                   *BankModel     `hexya:"many2one=ID"`       //":     fields.Many2One{RelationModel: h.Bank()},
	BankName               string         `hexya:"related=Bank.Name"` //": fields.Char{Related: "Bank.Name"},
	BankBIC                string         `hexya:"related=Bank.BIC"`  //":  fields.Char{Related: "Bank.BIC"},
	Sequence               int            //": fields.Integer{},
	Currency               CurrencyModel  `hexya:"many2one=ID"` //": fields.Many2One{RelationModel: h.Currency()},
	Company                []CompanyModel `hexya:"many2one=ID"` //": fields.Many2One{RelationModel: h.Company(), Required: true, Default: func(env models.Environment) interface{} {
}
