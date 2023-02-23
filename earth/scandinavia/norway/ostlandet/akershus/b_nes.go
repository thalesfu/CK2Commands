package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内斯NesBarony struct {
	feud.BaseBarony
}

var BNes内斯 feud.Barony = &内斯NesBarony{}

func init() {
	f := BNes内斯.(*内斯NesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nes",
		TitleName: "内斯",
		TitleCode: "b_nes",
	}
}
