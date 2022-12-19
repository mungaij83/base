package models

type PartnerIndustryModel struct {
	Name     string `hexya:"translated"`
	FullName string `hexya:"translated"`
	Active   bool   `hexya:"default=true"`
}
