package models

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/loader"
)

type LangModel struct {
	Name         string `hexya:"required;unique"`
	Code         string `hexya:"display_name:Locale Code;required;help=This field is used to set/get locales for user;unique"`
	ISOCode      string `hexya:"help=This ISO code is the name of PO files to use for translations"`
	Translatable bool
	Active       bool
	Direction    string `hexya:"options=ltr:Left-to-Right,rtl:Right-to-left;required;default=ltr"`
	DateFormat   string `hexya:"required;default=2006-01-02"`
	TimeFormat   string `hexya:"required;default=15:04:05"`
	Grouping     string `hexya:"display_name=Separator Format;required;default=[];help=The Separator Format should be like [,n] where 0 < n :starting from Unit digit.\"-1 will end the separation. e.g. [3,2,-1] will represent 106500 to be 1,06,500\"[1,2,-1] will represent it to be 106,50,0;[3] will represent it as 106,500.\"Provided \',\' as the thousand separator in each case."`
	DecimalPoint string `hexya:"display_name=Decimal Separator;required;default=."`
	ThousandsSep string `hexya:"display_name=Thousands Separator;default=\",\""`
}

type TranslationModel struct {
}


func InitTranslations() {
	loader.NewTypedModel(LangModel{})
	loader.NewTypedModel(TranslationModel{})
}