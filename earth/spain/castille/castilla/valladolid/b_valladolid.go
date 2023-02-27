package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴利亚多利德ValladolidBarony struct {
	feud.BaseBarony
}

var BValladolid巴利亚多利德 feud.Barony = &巴利亚多利德ValladolidBarony{}

func init() {
    f := BValladolid巴利亚多利德.(*巴利亚多利德ValladolidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valladolid",
		TitleName: "巴利亚多利德",
		TitleCode: "b_valladolid",
	}
}
