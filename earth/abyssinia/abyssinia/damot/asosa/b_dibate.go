package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪巴特DibateBarony struct {
	feud.BaseBarony
}

var BDibate迪巴特 feud.Barony = &迪巴特DibateBarony{}

func init() {
	f := BDibate迪巴特.(*迪巴特DibateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dibate",
		TitleName: "迪巴特",
		TitleCode: "b_dibate",
	}
}
