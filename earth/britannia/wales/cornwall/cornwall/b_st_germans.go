package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣杰曼斯St_germansBarony struct {
	feud.BaseBarony
}

var BSt_germans圣杰曼斯 feud.Barony = &圣杰曼斯St_germansBarony{}

func init() {
    f := BSt_germans圣杰曼斯.(*圣杰曼斯St_germansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_germans",
		TitleName: "圣杰曼斯",
		TitleCode: "b_st_germans",
	}
}
