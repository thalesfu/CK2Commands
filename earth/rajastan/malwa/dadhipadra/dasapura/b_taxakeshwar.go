package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德叉计湿伐罗TaxakeshwarBarony struct {
	feud.BaseBarony
}

var BTaxakeshwar德叉计湿伐罗 feud.Barony = &德叉计湿伐罗TaxakeshwarBarony{}

func init() {
	f := BTaxakeshwar德叉计湿伐罗.(*德叉计湿伐罗TaxakeshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taxakeshwar",
		TitleName: "德叉计湿伐罗",
		TitleCode: "b_taxakeshwar",
	}
}
