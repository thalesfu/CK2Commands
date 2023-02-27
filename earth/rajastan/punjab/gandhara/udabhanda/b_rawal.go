package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瓦尔RawalBarony struct {
	feud.BaseBarony
}

var BRawal拉瓦尔 feud.Barony = &拉瓦尔RawalBarony{}

func init() {
    f := BRawal拉瓦尔.(*拉瓦尔RawalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rawal",
		TitleName: "拉瓦尔",
		TitleCode: "b_rawal",
	}
}
