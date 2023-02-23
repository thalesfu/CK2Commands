package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔毛DalmauBarony struct {
	feud.BaseBarony
}

var BDalmau达尔毛 feud.Barony = &达尔毛DalmauBarony{}

func init() {
	f := BDalmau达尔毛.(*达尔毛DalmauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalmau",
		TitleName: "达尔毛",
		TitleCode: "b_dalmau",
	}
}
