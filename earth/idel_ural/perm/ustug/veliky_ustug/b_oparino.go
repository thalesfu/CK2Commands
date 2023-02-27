package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥帕里诺OparinoBarony struct {
	feud.BaseBarony
}

var BOparino奥帕里诺 feud.Barony = &奥帕里诺OparinoBarony{}

func init() {
    f := BOparino奥帕里诺.(*奥帕里诺OparinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oparino",
		TitleName: "奥帕里诺",
		TitleCode: "b_oparino",
	}
}
