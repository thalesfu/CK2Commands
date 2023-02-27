package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃森Essen_bergBarony struct {
	feud.BaseBarony
}

var BEssen_berg埃森 feud.Barony = &埃森Essen_bergBarony{}

func init() {
    f := BEssen_berg埃森.(*埃森Essen_bergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "essen_berg",
		TitleName: "埃森",
		TitleCode: "b_essen_berg",
	}
}
