package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁万杰马Ruwan_jemaBarony struct {
	feud.BaseBarony
}

var BRuwan_jema鲁万杰马 feud.Barony = &鲁万杰马Ruwan_jemaBarony{}

func init() {
    f := BRuwan_jema鲁万杰马.(*鲁万杰马Ruwan_jemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruwan_jema",
		TitleName: "鲁万杰马",
		TitleCode: "b_ruwan_jema",
	}
}
