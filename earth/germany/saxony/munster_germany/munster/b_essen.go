package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃森EssenBarony struct {
	feud.BaseBarony
}

var BEssen埃森 feud.Barony = &埃森EssenBarony{}

func init() {
    f := BEssen埃森.(*埃森EssenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "essen",
		TitleName: "埃森",
		TitleCode: "b_essen",
	}
}
