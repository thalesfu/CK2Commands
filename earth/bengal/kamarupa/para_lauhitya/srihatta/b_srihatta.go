package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利诃吒SrihattaBarony struct {
	feud.BaseBarony
}

var BSrihatta室利诃吒 feud.Barony = &室利诃吒SrihattaBarony{}

func init() {
    f := BSrihatta室利诃吒.(*室利诃吒SrihattaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srihatta",
		TitleName: "室利诃吒",
		TitleCode: "b_srihatta",
	}
}
