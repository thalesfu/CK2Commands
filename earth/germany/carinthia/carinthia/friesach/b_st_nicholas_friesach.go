package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣尼古拉斯St_nicholas_friesachBarony struct {
	feud.BaseBarony
}

var BSt_nicholas_friesach圣尼古拉斯 feud.Barony = &圣尼古拉斯St_nicholas_friesachBarony{}

func init() {
    f := BSt_nicholas_friesach圣尼古拉斯.(*圣尼古拉斯St_nicholas_friesachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_nicholas_friesach",
		TitleName: "圣尼古拉斯",
		TitleCode: "b_st_nicholas_friesach",
	}
}
