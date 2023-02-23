package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀喇沙尔KarasharBarony struct {
	feud.BaseBarony
}

var BKarashar喀喇沙尔 feud.Barony = &喀喇沙尔KarasharBarony{}

func init() {
	f := BKarashar喀喇沙尔.(*喀喇沙尔KarasharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karashar",
		TitleName: "喀喇沙尔",
		TitleCode: "b_karashar",
	}
}
