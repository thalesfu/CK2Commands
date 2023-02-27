package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪弗当日DifferdangeBarony struct {
	feud.BaseBarony
}

var BDifferdange迪弗当日 feud.Barony = &迪弗当日DifferdangeBarony{}

func init() {
    f := BDifferdange迪弗当日.(*迪弗当日DifferdangeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "differdange",
		TitleName: "迪弗当日",
		TitleCode: "b_differdange",
	}
}
