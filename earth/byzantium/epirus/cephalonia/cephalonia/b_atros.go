package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特罗斯AtrosBarony struct {
	feud.BaseBarony
}

var BAtros阿特罗斯 feud.Barony = &阿特罗斯AtrosBarony{}

func init() {
    f := BAtros阿特罗斯.(*阿特罗斯AtrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atros",
		TitleName: "阿特罗斯",
		TitleCode: "b_atros",
	}
}
