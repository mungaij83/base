package models

import (
	"github.com/hexya-erp/hexya/src/models/loader"
	"time"
)

type UserLogModel struct {
}

func (UserLogModel) OrderFields() []string {
	return []string{"id"}
}

type UserModel struct {
	Partner        *PartnerModel  `hexya:"display_name=Related Partner;many2one=id;embed;help=Partner-related data of the user"`
	Login          string         `hexya:"required;unique;help=Used to log into the system"`
	Password       string         `hexya:"noCopy;stored;help=Keep empty if you don't want the user to be able to connect on the system."`
	NewPassword    string         `hexya:"display_name=Set Password;stored=false;help=Specify a value only when creating a user or if you'rechanging the user's password, otherwise leave empty. Aftera change of password, the user has to login again."`
	Signature      string         `hexya:"type=text;display_name=Email Signature"`
	Active         bool           `hexya:"default=true"`
	ActivePartner  bool           `hexya:"default=true;readOnly;display_name=Partner Is Active"`
	ActionID       string         `hexya:"display_name=Home Action;help=If specified, this action will be opened at log on for this user, in addition to the standard menu."`
	Groups         []GroupModel   `json:"group_ids" hexya:"many2many=id"`
	Logs           []UserLogModel `json:"log_ids" hexya:"one2many=CreateUID;display_name=User Log Entries"`
	LoginDate      time.Time      `hexya:"type=datetime;Related=Logs.CreateDate;display_name=Latest authentication"`
	Share          bool           `hexya:"display_name=Share User;depends=Groups;type=Compute;Stored=true;help=External user with limited access, created only for the purpose of sharing data."`
	CompaniesCount int            `hexya:"display_name=Number of Companies;goType=int;type=compute"`
	TZOffset       string         `hexya:"display_name=Timezone Offset;type=compute"`
	Company        *CompanyModel  `hexya:"required;help=The company this user is currently working for."`
	Companies      []CompanyModel `json:"company_ids" hexya:"type=many2many=Id;"`
	GroupsCount    int            `hexya:"display_name:# Groups;type=compute;goType=int;help=Number of groups that apply to the current user"`
}

func InitUsers() {
	loader.NewTypedModel(GroupModel{})
	loader.NewTypedModel(UserModel{})
	loader.NewTypedModel(PartnerCategoryModel{})
	loader.NewTypedModel(PartnerTitleModel{})
	loader.NewTypedModel(PartnerIndustryModel{})
	loader.NewTypedModel(PartnerModel{})
	loader.NewTypedModel(CompanyModel{})
}
