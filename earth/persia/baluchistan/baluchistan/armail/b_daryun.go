package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达龙DaryunBarony struct {
	feud.BaseBarony
}

var BDaryun达龙 feud.Barony = &达龙DaryunBarony{}

func init() {
	f := BDaryun达龙.(*达龙DaryunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daryun",
		TitleName: "达龙",
		TitleCode: "b_daryun",
	}
}
