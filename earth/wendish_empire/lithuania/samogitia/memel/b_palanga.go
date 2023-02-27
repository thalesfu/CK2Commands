package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕兰加PalangaBarony struct {
	feud.BaseBarony
}

var BPalanga帕兰加 feud.Barony = &帕兰加PalangaBarony{}

func init() {
    f := BPalanga帕兰加.(*帕兰加PalangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palanga",
		TitleName: "帕兰加",
		TitleCode: "b_palanga",
	}
}
