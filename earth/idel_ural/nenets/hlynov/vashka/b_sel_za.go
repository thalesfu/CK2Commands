package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢利扎Sel_zaBarony struct {
	feud.BaseBarony
}

var BSel_za谢利扎 feud.Barony = &谢利扎Sel_zaBarony{}

func init() {
    f := BSel_za谢利扎.(*谢利扎Sel_zaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sel_za",
		TitleName: "谢利扎",
		TitleCode: "b_sel_za",
	}
}
