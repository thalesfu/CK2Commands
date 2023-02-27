package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁胡CruachuBarony struct {
	feud.BaseBarony
}

var BCruachu克鲁胡 feud.Barony = &克鲁胡CruachuBarony{}

func init() {
    f := BCruachu克鲁胡.(*克鲁胡CruachuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cruachu",
		TitleName: "克鲁胡",
		TitleCode: "b_cruachu",
	}
}
