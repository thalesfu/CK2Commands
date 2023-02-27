package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣大卫St_davidsBarony struct {
	feud.BaseBarony
}

var BSt_davids圣大卫 feud.Barony = &圣大卫St_davidsBarony{}

func init() {
    f := BSt_davids圣大卫.(*圣大卫St_davidsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_davids",
		TitleName: "圣大卫",
		TitleCode: "b_st_davids",
	}
}
