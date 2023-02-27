package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣皮埃尔St_pierreBarony struct {
	feud.BaseBarony
}

var BSt_pierre圣皮埃尔 feud.Barony = &圣皮埃尔St_pierreBarony{}

func init() {
    f := BSt_pierre圣皮埃尔.(*圣皮埃尔St_pierreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_pierre",
		TitleName: "圣皮埃尔",
		TitleCode: "b_st_pierre",
	}
}
