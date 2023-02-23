package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米岩鲁汉MiyanrudhanBarony struct {
	feud.BaseBarony
}

var BMiyanrudhan米岩鲁汉 feud.Barony = &米岩鲁汉MiyanrudhanBarony{}

func init() {
	f := BMiyanrudhan米岩鲁汉.(*米岩鲁汉MiyanrudhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miyanrudhan",
		TitleName: "米岩鲁汉",
		TitleCode: "b_miyanrudhan",
	}
}
