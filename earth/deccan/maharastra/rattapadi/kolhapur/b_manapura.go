package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马纳普拉ManapuraBarony struct {
	feud.BaseBarony
}

var BManapura马纳普拉 feud.Barony = &马纳普拉ManapuraBarony{}

func init() {
    f := BManapura马纳普拉.(*马纳普拉ManapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manapura",
		TitleName: "马纳普拉",
		TitleCode: "b_manapura",
	}
}
