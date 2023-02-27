package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢布内LubnyBarony struct {
	feud.BaseBarony
}

var BLubny卢布内 feud.Barony = &卢布内LubnyBarony{}

func init() {
    f := BLubny卢布内.(*卢布内LubnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lubny",
		TitleName: "卢布内",
		TitleCode: "b_lubny",
	}
}
