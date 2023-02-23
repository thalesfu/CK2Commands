package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉马朗伊斯GuimaraesBarony struct {
	feud.BaseBarony
}

var BGuimaraes吉马朗伊斯 feud.Barony = &吉马朗伊斯GuimaraesBarony{}

func init() {
	f := BGuimaraes吉马朗伊斯.(*吉马朗伊斯GuimaraesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guimaraes",
		TitleName: "吉马朗伊斯",
		TitleCode: "b_guimaraes",
	}
}
