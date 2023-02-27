package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈迪德Al_hadidBarony struct {
	feud.BaseBarony
}

var BAl_hadid哈迪德 feud.Barony = &哈迪德Al_hadidBarony{}

func init() {
    f := BAl_hadid哈迪德.(*哈迪德Al_hadidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_hadid",
		TitleName: "哈迪德",
		TitleCode: "b_al_hadid",
	}
}
