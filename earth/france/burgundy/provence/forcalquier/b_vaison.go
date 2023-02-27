package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦松VaisonBarony struct {
	feud.BaseBarony
}

var BVaison韦松 feud.Barony = &韦松VaisonBarony{}

func init() {
    f := BVaison韦松.(*韦松VaisonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaison",
		TitleName: "韦松",
		TitleCode: "b_vaison",
	}
}
