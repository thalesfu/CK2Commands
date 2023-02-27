package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬乔沃PancsovaBarony struct {
	feud.BaseBarony
}

var BPancsova蓬乔沃 feud.Barony = &蓬乔沃PancsovaBarony{}

func init() {
    f := BPancsova蓬乔沃.(*蓬乔沃PancsovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pancsova",
		TitleName: "蓬乔沃",
		TitleCode: "b_pancsova",
	}
}
