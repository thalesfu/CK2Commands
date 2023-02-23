package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内索斯NesosBarony struct {
	feud.BaseBarony
}

var BNesos内索斯 feud.Barony = &内索斯NesosBarony{}

func init() {
	f := BNesos内索斯.(*内索斯NesosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nesos",
		TitleName: "内索斯",
		TitleCode: "b_nesos",
	}
}
