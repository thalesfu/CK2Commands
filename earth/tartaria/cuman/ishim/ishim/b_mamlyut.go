package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马姆柳特MamlyutBarony struct {
	feud.BaseBarony
}

var BMamlyut马姆柳特 feud.Barony = &马姆柳特MamlyutBarony{}

func init() {
    f := BMamlyut马姆柳特.(*马姆柳特MamlyutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamlyut",
		TitleName: "马姆柳特",
		TitleCode: "b_mamlyut",
	}
}
