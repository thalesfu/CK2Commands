package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费马耶FemayeBarony struct {
	feud.BaseBarony
}

var BFemaye费马耶 feud.Barony = &费马耶FemayeBarony{}

func init() {
    f := BFemaye费马耶.(*费马耶FemayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "femaye",
		TitleName: "费马耶",
		TitleCode: "b_femaye",
	}
}
