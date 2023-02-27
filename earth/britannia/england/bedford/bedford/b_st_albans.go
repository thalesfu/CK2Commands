package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣奥尔本斯St_albansBarony struct {
	feud.BaseBarony
}

var BSt_albans圣奥尔本斯 feud.Barony = &圣奥尔本斯St_albansBarony{}

func init() {
    f := BSt_albans圣奥尔本斯.(*圣奥尔本斯St_albansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_albans",
		TitleName: "圣奥尔本斯",
		TitleCode: "b_st_albans",
	}
}
