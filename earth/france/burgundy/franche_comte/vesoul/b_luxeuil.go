package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕克瑟伊LuxeuilBarony struct {
	feud.BaseBarony
}

var BLuxeuil吕克瑟伊 feud.Barony = &吕克瑟伊LuxeuilBarony{}

func init() {
    f := BLuxeuil吕克瑟伊.(*吕克瑟伊LuxeuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luxeuil",
		TitleName: "吕克瑟伊",
		TitleCode: "b_luxeuil",
	}
}
