package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈梅拉LagomeraBarony struct {
	feud.BaseBarony
}

var BLagomera戈梅拉 feud.Barony = &戈梅拉LagomeraBarony{}

func init() {
	f := BLagomera戈梅拉.(*戈梅拉LagomeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagomera",
		TitleName: "戈梅拉",
		TitleCode: "b_lagomera",
	}
}
