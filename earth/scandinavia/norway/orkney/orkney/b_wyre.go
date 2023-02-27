package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 怀尔WyreBarony struct {
	feud.BaseBarony
}

var BWyre怀尔 feud.Barony = &怀尔WyreBarony{}

func init() {
    f := BWyre怀尔.(*怀尔WyreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wyre",
		TitleName: "怀尔",
		TitleCode: "b_wyre",
	}
}
