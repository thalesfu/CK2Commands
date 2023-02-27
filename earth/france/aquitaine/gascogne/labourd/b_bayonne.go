package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴约讷BayonneBarony struct {
	feud.BaseBarony
}

var BBayonne巴约讷 feud.Barony = &巴约讷BayonneBarony{}

func init() {
    f := BBayonne巴约讷.(*巴约讷BayonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayonne",
		TitleName: "巴约讷",
		TitleCode: "b_bayonne",
	}
}
