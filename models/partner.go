package models

import (
	"time"
)

type PartnerTitleModel struct {
	Name     string //   fields.Char{String: "Title", Required: true, Translate: true, Unique: true},
	Shortcut string //": fields.Char{String: "Abbreviation", Translate: true},
}

type PartnerCategoryModel struct {
	Name     string                 //:  fields.Char{String: "Tag Name", Required: true, Translate: true},
	Color    int                    //": fields.Integer{String: "Color Index"},
	Parent   *PartnerCategoryModel  // fields.Many2One{RelationModel: h.PartnerCategory(), String: "Parent Tag", Index: true, OnDelete: models.Cascade},
	Children []PartnerCategoryModel //": fields.One2Many{RelationModel: h.PartnerCategory(), Copy: true, ReverseFK: "Parent", String: "Children Tags"},
	Active   bool                   // fields.Boolean{Default: models.DefaultValue(true), Required: true, Help: "The active field allows you to hide the category without removing it."},
	Partners []PartnerModel         //: fields.Many2Many{RelationModel: h.Partner()},
}

type PartnerModel struct {
	Name                  string               `hexya:"required;index;noCopy"`
	Date                  time.Time            `hexya:"index"`
	Title                 PartnerTitleModel    `hexya:"many2one"`
	Parent                *PartnerModel        `hexya:"many2one;index"`
	ParentName            string               `hexya:"related=Parent.Name;readOnly"`
	Children              []PartnerModel       `hexya:"one2many;index"`
	Ref                   string               `hexya:"display_name=Internal Reference;index"` //: fields.Char{String: "Internal Reference", Index: true},
	Lang                  string               `hexya:"display_name=Language;index"`
	ActiveLangCount       int                  //": fields.Integer{Compute: h.Partner().Methods().ComputeActiveLangCount(), GoType: new(int)},
	TZ                    string               //: fields.Char{ String: "Timezone", Default: func(env models.Environment) interface{} { return env.Context().GetString("tz")},
	TZOffset              string               //": fields.Char{ Compute: h.Partner().Methods().ComputeTZOffset(), String:  "Timezone Offset", Depends: []string{"TZ"}},
	User                  UserModel            //": fields.Many2One{ RelationModel: h.User(), String:        "Salesperson", Help: "The internal user that is in charge of communicating with this contact if any."},
	VAT                   string               //": fields.Char{String: "TIN", Help: `Tax Identification Number. Fill it if the company is subjected to taxes. Used by the some of the legal statements.`},
	SameVATPartner        *PartnerModel        //": fields.Many2One{String: "Partner with same Tax ID", RelationModel: h.Partner(), Compute:       h.Partner().Methods().ComputeSameVATPartner()},
	Banks                 BankAccountModel     //": fields.One2Many{ String: "Bank Accounts", RelationModel: h.BankAccount(), ReverseFK: "Partner"},
	Website               string               //": fields.Char{Help: "Website of Partner or Company"},
	Comment               string               `hexya:"display_name=Notes"`
	Categories            PartnerCategoryModel //": fields.Many2Many{ RelationModel: h.PartnerCategory(), String: "Tags", JSON: "category_ids", Default: func(env models.Environment) interface{} { return h.PartnerCategory().Browse(env, []int64{env.Context().GetInteger("category_id")})}},
	CreditLimit           float32
	Barcode               string
	Active                bool   `hexya:"required;default=trie"`
	Employee              bool   `hexya:"default=true;help=Check this box if this contact is an Employee."`
	Function              string `hexya:"display_name=Job Position"`
	Type                  string `hexya:"option=contact:Contact,invoice:Invoice Address,delivery:Shipping Address,other:Other Address,private:Private Address"`
	Street                string
	Street2               string
	Zip                   string
	City                  string
	State                 CountryStateModel //": fields.Many2One{RelationModel: h.CountryState(), Filter: q.CountryState().Country().EqualsEval("country_id"), OnDelete: models.Restrict},
	Country               CountryModel      //": fields.Many2One{RelationModel: h.Country(), OnDelete: models.Restrict, OnChangeFilters: h.Partner().Methods().OnchangeCountryFilters()},
	Latitude              float32           //":  fields.Float{String: "Geo Latitude", Digits: nbutils.Digits{Precision: 16, Scale: 5}},
	Longitude             float32           //": fields.Float{String: "Geo Longitude", Digits: nbutils.Digits{Precision: 16, Scale: 5}},
	Email                 string            //":     fields.Char{OnChange: h.Partner().Methods().OnchangeEmail()},
	EmailFormatted        string            //": fields.Char{Compute: h.Partner().Methods().ComputeEmailFormatted(), Help: "Formatted email address 'Name <email@domain>'", Depends: []string{"Name", "Email"}},
	Phone                 string
	Mobile                string
	IsCompany             bool                 `hexya:"default=false;help=Check if the contact is a company, otherwise it is a person"`
	Industry              PartnerIndustryModel //": fields.Many2One{RelationModel: h.PartnerIndustry()},
	CompanyType           string               `hexya:"option=person:Individual,company:Company"`
	Company               []CompanyModel   `hexya:"many2one=id"`      //": fields.Many2One{RelationModel: h.Company()},
	Color                 int
	Users                 UserModel    `json:"user_ids" hexya:"one2many=id"`
	PartnerShare          bool         `hexya:"display_name=Share Partner;Stored;Depends=Users,Users.Share"`
	ContactAddress        string       `hexya:"display_name=Complete Address;Depends:Street,Street2,Zip,City,State,Country,Country.AddressFormat,Country.Code,Country.Name,CompanyName,State.Code,State.Name"`
	CommercialPartner     *PartnerModel `hexya:"many2one;display_name=Commercial Entity;index;Stored;depends=IsCompany,Parent,Parent.CommercialPartner"`
	CommercialCompanyName string       // ": fields.Char{ Compute: h.Partner().Methods().ComputeCommercialCompanyName(), Stored: true, Depends: []string{"CompanyName", "Parent", "Parent.IsCompany", "CommercialPartner", "CommercialPartner.Name"}},
	CompanyName           string
	Image                 []byte `hexya:"help=This field holds the image used as avatar for this contact, limited to 1024x1024px"`
	ImageMedium           []byte `hexya:"help=Medium-sized image of this contact. It is automatically resized as a 128x128px image, with aspect ratio preserved. Use this field in form views or some kanban views."`
	ImageSmall            []byte `hexya:"help=Small-sized image of this contact. It is automatically resized as a 64x64px image, with aspect ratio preserved. Use this field anywhere a small image is required."`
}
