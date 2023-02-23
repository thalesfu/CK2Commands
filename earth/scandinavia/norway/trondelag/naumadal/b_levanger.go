package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱旺厄尔LevangerBarony struct {
	feud.BaseBarony
}

var BLevanger莱旺厄尔 feud.Barony = &莱旺厄尔LevangerBarony{}

func init() {
	f := BLevanger莱旺厄尔.(*莱旺厄尔LevangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "levanger",
		TitleName: "莱旺厄尔",
		TitleCode: "b_levanger",
	}
}
