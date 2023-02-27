package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富加El_foghaBarony struct {
	feud.BaseBarony
}

var BEl_fogha富加 feud.Barony = &富加El_foghaBarony{}

func init() {
    f := BEl_fogha富加.(*富加El_foghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_fogha",
		TitleName: "富加",
		TitleCode: "b_el_fogha",
	}
}
