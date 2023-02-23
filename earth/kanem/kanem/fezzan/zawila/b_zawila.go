package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰维莱ZawilaBarony struct {
	feud.BaseBarony
}

var BZawila宰维莱 feud.Barony = &宰维莱ZawilaBarony{}

func init() {
	f := BZawila宰维莱.(*宰维莱ZawilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zawila",
		TitleName: "宰维莱",
		TitleCode: "b_zawila",
	}
}
