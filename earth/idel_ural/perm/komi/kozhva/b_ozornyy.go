package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥佐尔内OzornyyBarony struct {
	feud.BaseBarony
}

var BOzornyy奥佐尔内 feud.Barony = &奥佐尔内OzornyyBarony{}

func init() {
    f := BOzornyy奥佐尔内.(*奥佐尔内OzornyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozornyy",
		TitleName: "奥佐尔内",
		TitleCode: "b_ozornyy",
	}
}
