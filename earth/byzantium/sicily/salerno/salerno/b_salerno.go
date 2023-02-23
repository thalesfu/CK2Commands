package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨莱诺SalernoBarony struct {
	feud.BaseBarony
}

var BSalerno萨莱诺 feud.Barony = &萨莱诺SalernoBarony{}

func init() {
	f := BSalerno萨莱诺.(*萨莱诺SalernoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salerno",
		TitleName: "萨莱诺",
		TitleCode: "b_salerno",
	}
}
