package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马克西米利安St_maximillianBarony struct {
	feud.BaseBarony
}

var BSt_maximillian圣马克西米利安 feud.Barony = &圣马克西米利安St_maximillianBarony{}

func init() {
    f := BSt_maximillian圣马克西米利安.(*圣马克西米利安St_maximillianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_maximillian",
		TitleName: "圣马克西米利安",
		TitleCode: "b_st_maximillian",
	}
}
