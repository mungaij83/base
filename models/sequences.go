package models

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/loader"
	"time"
)

type SequenceModel struct {
	Name string `hexya:"required"`
	Code string `hexya:"display_name=Sequence Code"`
	Implementation string `hexya:"type=selection;required;options=standard:Standard,no_gap:No Gap;default=standard;help=Two sequence object implementations are offered: Standard and 'No gap'.
The latter is slower than the former but forbids any
gap in the sequence (while they are possible in the former)."`
	Active           bool                   `hexya:"required;default=true"`
	Prefix           string                 `hexya:"help=Prefix value of the record for the sequence"`
	Suffix           string                 `hexya:"help=Suffix value of the record for the sequence"`
	NumberNext       int64                  `hexya:"required;display_name=Next Number;default=1;help=Next number of this sequence"`
	NumberNextActual int64                  `hexya:"type=compute;depends=NumberNext;display_name=Next Number;help=Next number that will be used. This number can be incremented frequently so the displayed value might already be obsolete"`
	NumberIncrement  int                    `hexya:"required;display_name=Step;default=1;help=The next number of the sequence will be incremented by this number"`
	Padding          int                    `hexya:"default=0;required;display_name=Sequence Size;help=Hexya will automatically adds some '0' on the left of the 'Next Number' to get the required padding size."`
	Company          *CompanyModel          `hexya:"many2one=id;default=true"`
	UseDateRange     bool                   `hexya:"display_name=Use subsequences per Date Range"`
	DateRanges       SequenceDateRangeModel `hexya:"one2many=Sequence;display_name=Subsequences"`
}

type SequenceDateRangeModel struct {
	DateFrom         time.Time      `hexya:"type=date;required"`
	DateTo           time.Time      `hexya:"type=date;required"`
	Sequence         *SequenceModel `hexya:"many2one=id;display_name=Main Sequence;required;onDelete=cascase"`
	NumberNext       int64          `hexya:"required;display_name=Next Number;help=Next number of this sequence;default=1"`
	NumberNextActual int64          `hexya:"depends=NumberNext;display_name=Next Number;type=compute;help=Next number that will be used. This number can be incremented frequently so the displayed value might already be obsolete"`
}


func InitSequences() {
	loader.NewTypedModel(SequenceModel{})
	loader.NewTypedModel(SequenceDateRangeModel{})
}