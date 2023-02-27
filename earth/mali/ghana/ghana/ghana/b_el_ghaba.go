package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加巴El_ghabaBarony struct {
	feud.BaseBarony
}

var BEl_ghaba加巴 feud.Barony = &加巴El_ghabaBarony{}

func init() {
    f := BEl_ghaba加巴.(*加巴El_ghabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_ghaba",
		TitleName: "加巴",
		TitleCode: "b_el_ghaba",
	}
}
