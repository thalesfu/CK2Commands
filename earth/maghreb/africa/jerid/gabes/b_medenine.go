package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德宁MedenineBarony struct {
	feud.BaseBarony
}

var BMedenine梅德宁 feud.Barony = &梅德宁MedenineBarony{}

func init() {
	f := BMedenine梅德宁.(*梅德宁MedenineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medenine",
		TitleName: "梅德宁",
		TitleCode: "b_medenine",
	}
}
