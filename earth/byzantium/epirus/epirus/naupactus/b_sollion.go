package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索利翁SollionBarony struct {
	feud.BaseBarony
}

var BSollion索利翁 feud.Barony = &索利翁SollionBarony{}

func init() {
    f := BSollion索利翁.(*索利翁SollionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sollion",
		TitleName: "索利翁",
		TitleCode: "b_sollion",
	}
}
