package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡恩SkagenBarony struct {
	feud.BaseBarony
}

var BSkagen斯卡恩 feud.Barony = &斯卡恩SkagenBarony{}

func init() {
	f := BSkagen斯卡恩.(*斯卡恩SkagenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skagen",
		TitleName: "斯卡恩",
		TitleCode: "b_skagen",
	}
}
