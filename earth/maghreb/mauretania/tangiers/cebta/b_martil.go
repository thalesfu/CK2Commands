package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈尔提勒MartilBarony struct {
	feud.BaseBarony
}

var BMartil迈尔提勒 feud.Barony = &迈尔提勒MartilBarony{}

func init() {
    f := BMartil迈尔提勒.(*迈尔提勒MartilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martil",
		TitleName: "迈尔提勒",
		TitleCode: "b_martil",
	}
}
