package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱尔LeerBarony struct {
	feud.BaseBarony
}

var BLeer莱尔 feud.Barony = &莱尔LeerBarony{}

func init() {
	f := BLeer莱尔.(*莱尔LeerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leer",
		TitleName: "莱尔",
		TitleCode: "b_leer",
	}
}
