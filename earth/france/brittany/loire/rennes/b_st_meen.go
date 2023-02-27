package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣梅昂St_meenBarony struct {
	feud.BaseBarony
}

var BSt_meen圣梅昂 feud.Barony = &圣梅昂St_meenBarony{}

func init() {
    f := BSt_meen圣梅昂.(*圣梅昂St_meenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_meen",
		TitleName: "圣梅昂",
		TitleCode: "b_st_meen",
	}
}
