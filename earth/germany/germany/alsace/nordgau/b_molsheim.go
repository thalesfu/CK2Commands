package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔塞姆MolsheimBarony struct {
	feud.BaseBarony
}

var BMolsheim莫尔塞姆 feud.Barony = &莫尔塞姆MolsheimBarony{}

func init() {
	f := BMolsheim莫尔塞姆.(*莫尔塞姆MolsheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molsheim",
		TitleName: "莫尔塞姆",
		TitleCode: "b_molsheim",
	}
}
