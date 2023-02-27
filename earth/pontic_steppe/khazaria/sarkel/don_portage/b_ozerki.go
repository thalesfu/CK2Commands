package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥泽尔基OzerkiBarony struct {
	feud.BaseBarony
}

var BOzerki奥泽尔基 feud.Barony = &奥泽尔基OzerkiBarony{}

func init() {
    f := BOzerki奥泽尔基.(*奥泽尔基OzerkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozerki",
		TitleName: "奥泽尔基",
		TitleCode: "b_ozerki",
	}
}
