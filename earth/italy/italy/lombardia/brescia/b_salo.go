package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨洛SaloBarony struct {
	feud.BaseBarony
}

var BSalo萨洛 feud.Barony = &萨洛SaloBarony{}

func init() {
    f := BSalo萨洛.(*萨洛SaloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salo",
		TitleName: "萨洛",
		TitleCode: "b_salo",
	}
}
