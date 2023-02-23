package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊克什基莱UxkullBarony struct {
	feud.BaseBarony
}

var BUxkull伊克什基莱 feud.Barony = &伊克什基莱UxkullBarony{}

func init() {
	f := BUxkull伊克什基莱.(*伊克什基莱UxkullBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uxkull",
		TitleName: "伊克什基莱",
		TitleCode: "b_uxkull",
	}
}
