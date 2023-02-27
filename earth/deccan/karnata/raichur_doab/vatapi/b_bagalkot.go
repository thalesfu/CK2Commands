package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴加尔科特BagalkotBarony struct {
	feud.BaseBarony
}

var BBagalkot巴加尔科特 feud.Barony = &巴加尔科特BagalkotBarony{}

func init() {
    f := BBagalkot巴加尔科特.(*巴加尔科特BagalkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagalkot",
		TitleName: "巴加尔科特",
		TitleCode: "b_bagalkot",
	}
}
