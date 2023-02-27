package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科特布斯CottbusBarony struct {
	feud.BaseBarony
}

var BCottbus科特布斯 feud.Barony = &科特布斯CottbusBarony{}

func init() {
    f := BCottbus科特布斯.(*科特布斯CottbusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cottbus",
		TitleName: "科特布斯",
		TitleCode: "b_cottbus",
	}
}
