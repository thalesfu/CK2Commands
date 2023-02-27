package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔喀马MarqayomaBarony struct {
	feud.BaseBarony
}

var BMarqayoma马尔喀马 feud.Barony = &马尔喀马MarqayomaBarony{}

func init() {
    f := BMarqayoma马尔喀马.(*马尔喀马MarqayomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marqayoma",
		TitleName: "马尔喀马",
		TitleCode: "b_marqayoma",
	}
}
