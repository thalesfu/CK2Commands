package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦毕试KapisaBarony struct {
	feud.BaseBarony
}

var BKapisa迦毕试 feud.Barony = &迦毕试KapisaBarony{}

func init() {
	f := BKapisa迦毕试.(*迦毕试KapisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapisa",
		TitleName: "迦毕试",
		TitleCode: "b_kapisa",
	}
}
