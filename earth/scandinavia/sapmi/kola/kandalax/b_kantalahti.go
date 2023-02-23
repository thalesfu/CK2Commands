package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎塔拉赫蒂KantalahtiBarony struct {
	feud.BaseBarony
}

var BKantalahti坎塔拉赫蒂 feud.Barony = &坎塔拉赫蒂KantalahtiBarony{}

func init() {
	f := BKantalahti坎塔拉赫蒂.(*坎塔拉赫蒂KantalahtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kantalahti",
		TitleName: "坎塔拉赫蒂",
		TitleCode: "b_kantalahti",
	}
}
