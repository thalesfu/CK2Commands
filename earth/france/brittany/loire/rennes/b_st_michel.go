package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣米歇尔St_michelBarony struct {
	feud.BaseBarony
}

var BSt_michel圣米歇尔 feud.Barony = &圣米歇尔St_michelBarony{}

func init() {
    f := BSt_michel圣米歇尔.(*圣米歇尔St_michelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_michel",
		TitleName: "圣米歇尔",
		TitleCode: "b_st_michel",
	}
}
