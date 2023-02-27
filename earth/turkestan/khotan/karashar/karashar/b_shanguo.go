package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 山国ShanguoBarony struct {
	feud.BaseBarony
}

var BShanguo山国 feud.Barony = &山国ShanguoBarony{}

func init() {
    f := BShanguo山国.(*山国ShanguoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shanguo",
		TitleName: "山国",
		TitleCode: "b_shanguo",
	}
}
