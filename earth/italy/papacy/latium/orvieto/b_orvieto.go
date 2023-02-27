package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔维耶托OrvietoBarony struct {
	feud.BaseBarony
}

var BOrvieto奥尔维耶托 feud.Barony = &奥尔维耶托OrvietoBarony{}

func init() {
    f := BOrvieto奥尔维耶托.(*奥尔维耶托OrvietoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orvieto",
		TitleName: "奥尔维耶托",
		TitleCode: "b_orvieto",
	}
}
