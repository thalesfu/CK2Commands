package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙彼利埃MontpellierBarony struct {
	feud.BaseBarony
}

var BMontpellier蒙彼利埃 feud.Barony = &蒙彼利埃MontpellierBarony{}

func init() {
    f := BMontpellier蒙彼利埃.(*蒙彼利埃MontpellierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montpellier",
		TitleName: "蒙彼利埃",
		TitleCode: "b_montpellier",
	}
}
