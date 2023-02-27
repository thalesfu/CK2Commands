package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢布什LubuszBarony struct {
	feud.BaseBarony
}

var BLubusz卢布什 feud.Barony = &卢布什LubuszBarony{}

func init() {
    f := BLubusz卢布什.(*卢布什LubuszBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lubusz",
		TitleName: "卢布什",
		TitleCode: "b_lubusz",
	}
}
