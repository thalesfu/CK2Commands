package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔丹MardanBarony struct {
	feud.BaseBarony
}

var BMardan马尔丹 feud.Barony = &马尔丹MardanBarony{}

func init() {
    f := BMardan马尔丹.(*马尔丹MardanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mardan",
		TitleName: "马尔丹",
		TitleCode: "b_mardan",
	}
}
