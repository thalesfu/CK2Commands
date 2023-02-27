package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索拉诺SoranoBarony struct {
	feud.BaseBarony
}

var BSorano索拉诺 feud.Barony = &索拉诺SoranoBarony{}

func init() {
    f := BSorano索拉诺.(*索拉诺SoranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorano",
		TitleName: "索拉诺",
		TitleCode: "b_sorano",
	}
}
