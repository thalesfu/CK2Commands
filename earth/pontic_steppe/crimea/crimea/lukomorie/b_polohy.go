package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波洛吉PolohyBarony struct {
	feud.BaseBarony
}

var BPolohy波洛吉 feud.Barony = &波洛吉PolohyBarony{}

func init() {
    f := BPolohy波洛吉.(*波洛吉PolohyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polohy",
		TitleName: "波洛吉",
		TitleCode: "b_polohy",
	}
}
