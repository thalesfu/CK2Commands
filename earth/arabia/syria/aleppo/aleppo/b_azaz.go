package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿扎兹AzazBarony struct {
	feud.BaseBarony
}

var BAzaz阿扎兹 feud.Barony = &阿扎兹AzazBarony{}

func init() {
	f := BAzaz阿扎兹.(*阿扎兹AzazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azaz",
		TitleName: "阿扎兹",
		TitleCode: "b_azaz",
	}
}
