package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎查GachaBarony struct {
	feud.BaseBarony
}

var BGacha嘎查 feud.Barony = &嘎查GachaBarony{}

func init() {
    f := BGacha嘎查.(*嘎查GachaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gacha",
		TitleName: "嘎查",
		TitleCode: "b_gacha",
	}
}
