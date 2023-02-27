package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛什LochesBarony struct {
	feud.BaseBarony
}

var BLoches洛什 feud.Barony = &洛什LochesBarony{}

func init() {
    f := BLoches洛什.(*洛什LochesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loches",
		TitleName: "洛什",
		TitleCode: "b_loches",
	}
}
