package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德拉BadraBarony struct {
	feud.BaseBarony
}

var BBadra巴德拉 feud.Barony = &巴德拉BadraBarony{}

func init() {
    f := BBadra巴德拉.(*巴德拉BadraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badra",
		TitleName: "巴德拉",
		TitleCode: "b_badra",
	}
}
