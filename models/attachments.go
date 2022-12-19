package models

import (
	"github.com/hexya-erp/hexya/src/models/loader"
)

type AttachmentModel struct {
	loader.Model
	Name         string       `hexya:"display_name=Attachment Name;required"`
	Description  string       `hexya:"type=text"`
	ResName      string       `hexya:"display_name=Resource Name;type=compute;stored;depends=ResModel;ResID"`
	ResModel     string       `hexya:"index;readOnly;display_name=Resource Model;help=The database object this attachment will be attached to"`
	ResField     string       `hexya:"index;readOnly;display_name=Resource Field;"`
	ResID        int64        `hexya:"display_name=Resource ID;help=The record id this is attached to;index;readOnly"`
	Company      CompanyModel `hexya:"many2many=id;"`
	Type         string       `hexya:"type=selection;options=binary:File,url:URL;help=You can either upload a file from your computer or copy/paste an internet link to your file."`
	URL          string       `hexya:"size=1024;index"`
	Public       bool         `hexya:"display_name=Is a public document"`
	AccessToken  string       `hexya:"size=512"`
	Datas        []byte       `hexya:"display_name=File Content;depends=StoreFname,DBDatas;type=compute"`
	DBDatas      string       `hexya:"display_name=Database Data"`
	StoreFname   string       `hexya:"display_name=Stored Filename"`
	FileSize     int          `hexya:"goType=int"`
	CheckSum     string       `hexya:"display_name=Checksum/SHA1;size=40;index;readOnly"`
	MimeType     string       `hexya:"readOnly"`
	IndexContent string       `hexya:"type=text;readOnly;display_name=Indexed Content"`
}

func InitAttachments() {
	attachment := loader.NewTypedModel(AttachmentModel{})
	definer := loader.NewModelSet[loader.Model{}](attachment)
}
