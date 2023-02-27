package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托木TimBarony struct {
	feud.BaseBarony
}

var BTim托木 feud.Barony = &托木TimBarony{}

func init() {
    f := BTim托木.(*托木TimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tim",
		TitleName: "托木",
		TitleCode: "b_tim",
	}
}
