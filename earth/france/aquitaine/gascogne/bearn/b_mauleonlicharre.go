package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫莱翁利沙尔MauleonlicharreBarony struct {
	feud.BaseBarony
}

var BMauleonlicharre莫莱翁利沙尔 feud.Barony = &莫莱翁利沙尔MauleonlicharreBarony{}

func init() {
	f := BMauleonlicharre莫莱翁利沙尔.(*莫莱翁利沙尔MauleonlicharreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mauleonlicharre",
		TitleName: "莫莱翁利沙尔",
		TitleCode: "b_mauleonlicharre",
	}
}
