package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮聪达PitsundaBarony struct {
	feud.BaseBarony
}

var BPitsunda皮聪达 feud.Barony = &皮聪达PitsundaBarony{}

func init() {
    f := BPitsunda皮聪达.(*皮聪达PitsundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitsunda",
		TitleName: "皮聪达",
		TitleCode: "b_pitsunda",
	}
}
