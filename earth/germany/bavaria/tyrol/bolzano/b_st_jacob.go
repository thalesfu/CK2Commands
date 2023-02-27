package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣雅各伯St_jacobBarony struct {
	feud.BaseBarony
}

var BSt_jacob圣雅各伯 feud.Barony = &圣雅各伯St_jacobBarony{}

func init() {
    f := BSt_jacob圣雅各伯.(*圣雅各伯St_jacobBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_jacob",
		TitleName: "圣雅各伯",
		TitleCode: "b_st_jacob",
	}
}
