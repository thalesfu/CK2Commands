package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥米什OmisBarony struct {
	feud.BaseBarony
}

var BOmis奥米什 feud.Barony = &奥米什OmisBarony{}

func init() {
    f := BOmis奥米什.(*奥米什OmisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omis",
		TitleName: "奥米什",
		TitleCode: "b_omis",
	}
}
