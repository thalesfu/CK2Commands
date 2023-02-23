package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 度摩军荼DomakondaBarony struct {
	feud.BaseBarony
}

var BDomakonda度摩军荼 feud.Barony = &度摩军荼DomakondaBarony{}

func init() {
	f := BDomakonda度摩军荼.(*度摩军荼DomakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domakonda",
		TitleName: "度摩军荼",
		TitleCode: "b_domakonda",
	}
}
