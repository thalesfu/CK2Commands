package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉穆尔KaramulBarony struct {
	feud.BaseBarony
}

var BKaramul卡拉穆尔 feud.Barony = &卡拉穆尔KaramulBarony{}

func init() {
    f := BKaramul卡拉穆尔.(*卡拉穆尔KaramulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karamul",
		TitleName: "卡拉穆尔",
		TitleCode: "b_karamul",
	}
}
