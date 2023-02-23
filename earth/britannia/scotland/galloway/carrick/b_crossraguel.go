package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科罗斯拉结尔CrossraguelBarony struct {
	feud.BaseBarony
}

var BCrossraguel科罗斯拉结尔 feud.Barony = &科罗斯拉结尔CrossraguelBarony{}

func init() {
	f := BCrossraguel科罗斯拉结尔.(*科罗斯拉结尔CrossraguelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crossraguel",
		TitleName: "科罗斯拉结尔",
		TitleCode: "b_crossraguel",
	}
}
