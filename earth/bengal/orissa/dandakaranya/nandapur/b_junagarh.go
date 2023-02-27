package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周那姞利呬JunagarhBarony struct {
	feud.BaseBarony
}

var BJunagarh周那姞利呬 feud.Barony = &周那姞利呬JunagarhBarony{}

func init() {
    f := BJunagarh周那姞利呬.(*周那姞利呬JunagarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "junagarh",
		TitleName: "周那姞利呬",
		TitleCode: "b_junagarh",
	}
}
