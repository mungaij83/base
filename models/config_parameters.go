package models

type ConfigParameterModel struct {
	Key    string       `hexya:"index;required;unique"`
	Value  string       `hexya:"required"`
	Groups []GroupModel `hexya:"many2many=id"`
}
