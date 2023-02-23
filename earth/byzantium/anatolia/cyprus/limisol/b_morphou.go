package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔富MorphouBarony struct {
	feud.BaseBarony
}

var BMorphou莫尔富 feud.Barony = &莫尔富MorphouBarony{}

func init() {
	f := BMorphou莫尔富.(*莫尔富MorphouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morphou",
		TitleName: "莫尔富",
		TitleCode: "b_morphou",
	}
}
