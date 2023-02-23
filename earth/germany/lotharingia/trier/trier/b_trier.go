package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里尔TrierBarony struct {
	feud.BaseBarony
}

var BTrier特里尔 feud.Barony = &特里尔TrierBarony{}

func init() {
	f := BTrier特里尔.(*特里尔TrierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trier",
		TitleName: "特里尔",
		TitleCode: "b_trier",
	}
}
