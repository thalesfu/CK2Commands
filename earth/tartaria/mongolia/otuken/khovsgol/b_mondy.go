package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙德MondyBarony struct {
	feud.BaseBarony
}

var BMondy蒙德 feud.Barony = &蒙德MondyBarony{}

func init() {
    f := BMondy蒙德.(*蒙德MondyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mondy",
		TitleName: "蒙德",
		TitleCode: "b_mondy",
	}
}
