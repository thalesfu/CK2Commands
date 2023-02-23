package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧格拉OuguelaBarony struct {
	feud.BaseBarony
}

var BOuguela欧格拉 feud.Barony = &欧格拉OuguelaBarony{}

func init() {
	f := BOuguela欧格拉.(*欧格拉OuguelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouguela",
		TitleName: "欧格拉",
		TitleCode: "b_ouguela",
	}
}
