package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪伦DurenBarony struct {
	feud.BaseBarony
}

var BDuren迪伦 feud.Barony = &迪伦DurenBarony{}

func init() {
    f := BDuren迪伦.(*迪伦DurenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duren",
		TitleName: "迪伦",
		TitleCode: "b_duren",
	}
}
