package models

import "github.com/hexya-erp/hexya/src/models/loader"

type BankModel struct {
	Name    string `hexya:"required"`
	Street  string
	Street2 string
	Zip     string
	City    string
	State   CountryStateModel
	Country CountryModel //": fields.Many2One{RelationModel: h.Country(), OnChange: h.Bank().Methods().OnchangeCountry()},
	Email   string
	Phone   string
	Active  bool
	BIC     string `hexya:"string=Bank Identifier Code;index;help=Sometimes called BIC or Swift."`
}

func (BankModel) TableName() string {
	return "bank"
}


func InitBank() {
	loader.NewTypedModel(BankModel{})
	loader.NewTypedModel(BankAccountModel{})
}