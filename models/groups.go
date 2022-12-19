package models

type GroupModel struct {
	GroupID string `hexya:"required"`
	Name    string `hexya:"required;translate"`
}
