package models

import "github.com/hexya-erp/hexya/src/models/loader"

type ConfigSettingsTransientModel struct {
}

func InitSettings() {
	loader.NewTypedModel(AutoVacuumManualModel{})
	loader.NewTypedModel(ConfigSettingsTransientModel{})
	loader.NewTypedModel(ConfigParameterModel{})

	// Country specific
	loader.NewTypedModel(CurrencyModel{})
	loader.NewTypedModel(CurrencyRateModel{})
	loader.NewTypedModel(CountryGroupModel{})
	loader.NewTypedModel(CountryModel{})
	loader.NewTypedModel(CountryStateModel{})
}
