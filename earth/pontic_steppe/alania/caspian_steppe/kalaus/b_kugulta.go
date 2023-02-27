package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库古利塔KugultaBarony struct {
	feud.BaseBarony
}

var BKugulta库古利塔 feud.Barony = &库古利塔KugultaBarony{}

func init() {
    f := BKugulta库古利塔.(*库古利塔KugultaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kugulta",
		TitleName: "库古利塔",
		TitleCode: "b_kugulta",
	}
}
