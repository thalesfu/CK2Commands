package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 须南SunamBarony struct {
	feud.BaseBarony
}

var BSunam须南 feud.Barony = &须南SunamBarony{}

func init() {
	f := BSunam须南.(*须南SunamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sunam",
		TitleName: "须南",
		TitleCode: "b_sunam",
	}
}
