package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅希利布日MysliborzBarony struct {
	feud.BaseBarony
}

var BMysliborz梅希利布日 feud.Barony = &梅希利布日MysliborzBarony{}

func init() {
    f := BMysliborz梅希利布日.(*梅希利布日MysliborzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mysliborz",
		TitleName: "梅希利布日",
		TitleCode: "b_mysliborz",
	}
}
