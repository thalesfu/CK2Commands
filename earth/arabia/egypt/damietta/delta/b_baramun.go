package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉穆恩BaramunBarony struct {
	feud.BaseBarony
}

var BBaramun巴拉穆恩 feud.Barony = &巴拉穆恩BaramunBarony{}

func init() {
    f := BBaramun巴拉穆恩.(*巴拉穆恩BaramunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baramun",
		TitleName: "巴拉穆恩",
		TitleCode: "b_baramun",
	}
}
