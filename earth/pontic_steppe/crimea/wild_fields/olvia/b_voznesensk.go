package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃兹涅先斯克VoznesenskBarony struct {
	feud.BaseBarony
}

var BVoznesensk沃兹涅先斯克 feud.Barony = &沃兹涅先斯克VoznesenskBarony{}

func init() {
    f := BVoznesensk沃兹涅先斯克.(*沃兹涅先斯克VoznesenskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voznesensk",
		TitleName: "沃兹涅先斯克",
		TitleCode: "b_voznesensk",
	}
}
