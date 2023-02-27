package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科贡KogonBarony struct {
	feud.BaseBarony
}

var BKogon科贡 feud.Barony = &科贡KogonBarony{}

func init() {
    f := BKogon科贡.(*科贡KogonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kogon",
		TitleName: "科贡",
		TitleCode: "b_kogon",
	}
}
