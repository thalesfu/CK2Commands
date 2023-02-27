package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴什马奇卡BaszmackaBarony struct {
	feud.BaseBarony
}

var BBaszmacka巴什马奇卡 feud.Barony = &巴什马奇卡BaszmackaBarony{}

func init() {
    f := BBaszmacka巴什马奇卡.(*巴什马奇卡BaszmackaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baszmacka",
		TitleName: "巴什马奇卡",
		TitleCode: "b_baszmacka",
	}
}
