package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科都尼CordueneBarony struct {
	feud.BaseBarony
}

var BCorduene科都尼 feud.Barony = &科都尼CordueneBarony{}

func init() {
    f := BCorduene科都尼.(*科都尼CordueneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corduene",
		TitleName: "科都尼",
		TitleCode: "b_corduene",
	}
}
