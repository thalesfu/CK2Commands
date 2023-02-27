package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代明DemminBarony struct {
	feud.BaseBarony
}

var BDemmin代明 feud.Barony = &代明DemminBarony{}

func init() {
    f := BDemmin代明.(*代明DemminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demmin",
		TitleName: "代明",
		TitleCode: "b_demmin",
	}
}
