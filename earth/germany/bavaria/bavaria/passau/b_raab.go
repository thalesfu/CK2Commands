package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉布RaabBarony struct {
	feud.BaseBarony
}

var BRaab拉布 feud.Barony = &拉布RaabBarony{}

func init() {
    f := BRaab拉布.(*拉布RaabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raab",
		TitleName: "拉布",
		TitleCode: "b_raab",
	}
}
