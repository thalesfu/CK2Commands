package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达卡DakkarBarony struct {
	feud.BaseBarony
}

var BDakkar达卡 feud.Barony = &达卡DakkarBarony{}

func init() {
	f := BDakkar达卡.(*达卡DakkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dakkar",
		TitleName: "达卡",
		TitleCode: "b_dakkar",
	}
}
