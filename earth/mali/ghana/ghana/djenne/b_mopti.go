package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫普提MoptiBarony struct {
	feud.BaseBarony
}

var BMopti莫普提 feud.Barony = &莫普提MoptiBarony{}

func init() {
    f := BMopti莫普提.(*莫普提MoptiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mopti",
		TitleName: "莫普提",
		TitleCode: "b_mopti",
	}
}
