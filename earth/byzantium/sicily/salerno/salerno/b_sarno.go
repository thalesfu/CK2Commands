package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔诺SarnoBarony struct {
	feud.BaseBarony
}

var BSarno萨尔诺 feud.Barony = &萨尔诺SarnoBarony{}

func init() {
	f := BSarno萨尔诺.(*萨尔诺SarnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarno",
		TitleName: "萨尔诺",
		TitleCode: "b_sarno",
	}
}
