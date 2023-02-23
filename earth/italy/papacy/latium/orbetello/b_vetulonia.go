package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦图罗尼亚VetuloniaBarony struct {
	feud.BaseBarony
}

var BVetulonia韦图罗尼亚 feud.Barony = &韦图罗尼亚VetuloniaBarony{}

func init() {
	f := BVetulonia韦图罗尼亚.(*韦图罗尼亚VetuloniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vetulonia",
		TitleName: "韦图罗尼亚",
		TitleCode: "b_vetulonia",
	}
}
