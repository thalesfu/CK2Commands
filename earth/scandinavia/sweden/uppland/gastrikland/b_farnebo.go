package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔内布FarneboBarony struct {
	feud.BaseBarony
}

var BFarnebo费尔内布 feud.Barony = &费尔内布FarneboBarony{}

func init() {
    f := BFarnebo费尔内布.(*费尔内布FarneboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farnebo",
		TitleName: "费尔内布",
		TitleCode: "b_farnebo",
	}
}
