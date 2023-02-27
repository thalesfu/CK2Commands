package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈斯鲁格MasruqBarony struct {
	feud.BaseBarony
}

var BMasruq迈斯鲁格 feud.Barony = &迈斯鲁格MasruqBarony{}

func init() {
    f := BMasruq迈斯鲁格.(*迈斯鲁格MasruqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masruq",
		TitleName: "迈斯鲁格",
		TitleCode: "b_masruq",
	}
}
