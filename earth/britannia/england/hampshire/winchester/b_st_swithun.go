package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣斯威辛St_swithunBarony struct {
	feud.BaseBarony
}

var BSt_swithun圣斯威辛 feud.Barony = &圣斯威辛St_swithunBarony{}

func init() {
    f := BSt_swithun圣斯威辛.(*圣斯威辛St_swithunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_swithun",
		TitleName: "圣斯威辛",
		TitleCode: "b_st_swithun",
	}
}
