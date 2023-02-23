package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉济MarazaBarony struct {
	feud.BaseBarony
}

var BMaraza马拉济 feud.Barony = &马拉济MarazaBarony{}

func init() {
	f := BMaraza马拉济.(*马拉济MarazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maraza",
		TitleName: "马拉济",
		TitleCode: "b_maraza",
	}
}
