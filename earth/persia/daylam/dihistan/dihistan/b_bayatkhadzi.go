package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴亚特_哈济BayatkhadziBarony struct {
	feud.BaseBarony
}

var BBayatkhadzi巴亚特_哈济 feud.Barony = &巴亚特_哈济BayatkhadziBarony{}

func init() {
    f := BBayatkhadzi巴亚特_哈济.(*巴亚特_哈济BayatkhadziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayatkhadzi",
		TitleName: "巴亚特_哈济",
		TitleCode: "b_bayatkhadzi",
	}
}
