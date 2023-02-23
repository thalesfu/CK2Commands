package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔特CorteBarony struct {
	feud.BaseBarony
}

var BCorte科尔特 feud.Barony = &科尔特CorteBarony{}

func init() {
	f := BCorte科尔特.(*科尔特CorteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corte",
		TitleName: "科尔特",
		TitleCode: "b_corte",
	}
}
