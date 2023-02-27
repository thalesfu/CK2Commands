package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷基RykiBarony struct {
	feud.BaseBarony
}

var BRyki雷基 feud.Barony = &雷基RykiBarony{}

func init() {
    f := BRyki雷基.(*雷基RykiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryki",
		TitleName: "雷基",
		TitleCode: "b_ryki",
	}
}
