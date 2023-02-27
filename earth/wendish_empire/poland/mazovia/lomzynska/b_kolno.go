package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔诺KolnoBarony struct {
	feud.BaseBarony
}

var BKolno科尔诺 feud.Barony = &科尔诺KolnoBarony{}

func init() {
    f := BKolno科尔诺.(*科尔诺KolnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolno",
		TitleName: "科尔诺",
		TitleCode: "b_kolno",
	}
}
