package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱夫卡斯LefkasBarony struct {
	feud.BaseBarony
}

var BLefkas莱夫卡斯 feud.Barony = &莱夫卡斯LefkasBarony{}

func init() {
	f := BLefkas莱夫卡斯.(*莱夫卡斯LefkasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lefkas",
		TitleName: "莱夫卡斯",
		TitleCode: "b_lefkas",
	}
}
