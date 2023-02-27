package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦梨耶尼KalyaniBarony struct {
	feud.BaseBarony
}

var BKalyani迦梨耶尼 feud.Barony = &迦梨耶尼KalyaniBarony{}

func init() {
    f := BKalyani迦梨耶尼.(*迦梨耶尼KalyaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalyani",
		TitleName: "迦梨耶尼",
		TitleCode: "b_kalyani",
	}
}
