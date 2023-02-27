package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马洛St_maloBarony struct {
	feud.BaseBarony
}

var BSt_malo圣马洛 feud.Barony = &圣马洛St_maloBarony{}

func init() {
    f := BSt_malo圣马洛.(*圣马洛St_maloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_malo",
		TitleName: "圣马洛",
		TitleCode: "b_st_malo",
	}
}
