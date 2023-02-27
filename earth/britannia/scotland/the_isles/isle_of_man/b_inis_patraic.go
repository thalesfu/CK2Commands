package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕特里克岛Inis_patraicBarony struct {
	feud.BaseBarony
}

var BInis_patraic帕特里克岛 feud.Barony = &帕特里克岛Inis_patraicBarony{}

func init() {
    f := BInis_patraic帕特里克岛.(*帕特里克岛Inis_patraicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inis_patraic",
		TitleName: "帕特里克岛",
		TitleCode: "b_inis_patraic",
	}
}
