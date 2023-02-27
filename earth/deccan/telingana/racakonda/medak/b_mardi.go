package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛地MardiBarony struct {
	feud.BaseBarony
}

var BMardi玛地 feud.Barony = &玛地MardiBarony{}

func init() {
    f := BMardi玛地.(*玛地MardiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mardi",
		TitleName: "玛地",
		TitleCode: "b_mardi",
	}
}
