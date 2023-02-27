package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣雅凯姆St_jacquemeBarony struct {
	feud.BaseBarony
}

var BSt_jacqueme圣雅凯姆 feud.Barony = &圣雅凯姆St_jacquemeBarony{}

func init() {
    f := BSt_jacqueme圣雅凯姆.(*圣雅凯姆St_jacquemeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_jacqueme",
		TitleName: "圣雅凯姆",
		TitleCode: "b_st_jacqueme",
	}
}
