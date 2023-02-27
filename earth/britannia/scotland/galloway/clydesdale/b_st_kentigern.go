package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣肯蒂格恩St_kentigernBarony struct {
	feud.BaseBarony
}

var BSt_kentigern圣肯蒂格恩 feud.Barony = &圣肯蒂格恩St_kentigernBarony{}

func init() {
    f := BSt_kentigern圣肯蒂格恩.(*圣肯蒂格恩St_kentigernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_kentigern",
		TitleName: "圣肯蒂格恩",
		TitleCode: "b_st_kentigern",
	}
}
