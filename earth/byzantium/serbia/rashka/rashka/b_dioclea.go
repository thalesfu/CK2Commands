package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪奥克勒亚DiocleaBarony struct {
	feud.BaseBarony
}

var BDioclea迪奥克勒亚 feud.Barony = &迪奥克勒亚DiocleaBarony{}

func init() {
	f := BDioclea迪奥克勒亚.(*迪奥克勒亚DiocleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dioclea",
		TitleName: "迪奥克勒亚",
		TitleCode: "b_dioclea",
	}
}
