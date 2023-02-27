package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米泰纳MitenaBarony struct {
	feud.BaseBarony
}

var BMitena米泰纳 feud.Barony = &米泰纳MitenaBarony{}

func init() {
    f := BMitena米泰纳.(*米泰纳MitenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mitena",
		TitleName: "米泰纳",
		TitleCode: "b_mitena",
	}
}
