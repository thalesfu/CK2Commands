package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣尼克拉斯St_niklaasBarony struct {
	feud.BaseBarony
}

var BSt_niklaas圣尼克拉斯 feud.Barony = &圣尼克拉斯St_niklaasBarony{}

func init() {
    f := BSt_niklaas圣尼克拉斯.(*圣尼克拉斯St_niklaasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_niklaas",
		TitleName: "圣尼克拉斯",
		TitleCode: "b_st_niklaas",
	}
}
