package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙雅XayarBarony struct {
	feud.BaseBarony
}

var BXayar沙雅 feud.Barony = &沙雅XayarBarony{}

func init() {
	f := BXayar沙雅.(*沙雅XayarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xayar",
		TitleName: "沙雅",
		TitleCode: "b_xayar",
	}
}
