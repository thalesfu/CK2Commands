package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍卜沃SabwaBarony struct {
	feud.BaseBarony
}

var BSabwa舍卜沃 feud.Barony = &舍卜沃SabwaBarony{}

func init() {
    f := BSabwa舍卜沃.(*舍卜沃SabwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabwa",
		TitleName: "舍卜沃",
		TitleCode: "b_sabwa",
	}
}
