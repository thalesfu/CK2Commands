package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶内GeneBarony struct {
	feud.BaseBarony
}

var BGene耶内 feud.Barony = &耶内GeneBarony{}

func init() {
    f := BGene耶内.(*耶内GeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gene",
		TitleName: "耶内",
		TitleCode: "b_gene",
	}
}
