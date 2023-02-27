package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣卡斯伯特St_cuthbertBarony struct {
	feud.BaseBarony
}

var BSt_cuthbert圣卡斯伯特 feud.Barony = &圣卡斯伯特St_cuthbertBarony{}

func init() {
    f := BSt_cuthbert圣卡斯伯特.(*圣卡斯伯特St_cuthbertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_cuthbert",
		TitleName: "圣卡斯伯特",
		TitleCode: "b_st_cuthbert",
	}
}
