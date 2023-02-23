package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布朗库堡CastelobrancoBarony struct {
	feud.BaseBarony
}

var BCastelobranco布朗库堡 feud.Barony = &布朗库堡CastelobrancoBarony{}

func init() {
	f := BCastelobranco布朗库堡.(*布朗库堡CastelobrancoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelobranco",
		TitleName: "布朗库堡",
		TitleCode: "b_castelobranco",
	}
}
