package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯加尔Al_masqarBarony struct {
	feud.BaseBarony
}

var BAl_masqar马斯加尔 feud.Barony = &马斯加尔Al_masqarBarony{}

func init() {
    f := BAl_masqar马斯加尔.(*马斯加尔Al_masqarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_masqar",
		TitleName: "马斯加尔",
		TitleCode: "b_al_masqar",
	}
}
