package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁德尔DinderBarony struct {
	feud.BaseBarony
}

var BDinder丁德尔 feud.Barony = &丁德尔DinderBarony{}

func init() {
	f := BDinder丁德尔.(*丁德尔DinderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinder",
		TitleName: "丁德尔",
		TitleCode: "b_dinder",
	}
}
