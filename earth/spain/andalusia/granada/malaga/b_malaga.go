package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉加MalagaBarony struct {
	feud.BaseBarony
}

var BMalaga马拉加 feud.Barony = &马拉加MalagaBarony{}

func init() {
    f := BMalaga马拉加.(*马拉加MalagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malaga",
		TitleName: "马拉加",
		TitleCode: "b_malaga",
	}
}
