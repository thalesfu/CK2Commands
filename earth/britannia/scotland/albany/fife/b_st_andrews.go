package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣安德鲁St_andrewsBarony struct {
	feud.BaseBarony
}

var BSt_andrews圣安德鲁 feud.Barony = &圣安德鲁St_andrewsBarony{}

func init() {
    f := BSt_andrews圣安德鲁.(*圣安德鲁St_andrewsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_andrews",
		TitleName: "圣安德鲁",
		TitleCode: "b_st_andrews",
	}
}
