package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努库斯NukusBarony struct {
	feud.BaseBarony
}

var BNukus努库斯 feud.Barony = &努库斯NukusBarony{}

func init() {
	f := BNukus努库斯.(*努库斯NukusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nukus",
		TitleName: "努库斯",
		TitleCode: "b_nukus",
	}
}
