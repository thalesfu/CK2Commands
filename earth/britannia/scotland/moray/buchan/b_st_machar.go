package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马哈尔St_macharBarony struct {
	feud.BaseBarony
}

var BSt_machar圣马哈尔 feud.Barony = &圣马哈尔St_macharBarony{}

func init() {
    f := BSt_machar圣马哈尔.(*圣马哈尔St_macharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_machar",
		TitleName: "圣马哈尔",
		TitleCode: "b_st_machar",
	}
}
