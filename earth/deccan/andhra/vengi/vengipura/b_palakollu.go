package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗科鲁PalakolluBarony struct {
	feud.BaseBarony
}

var BPalakollu波罗科鲁 feud.Barony = &波罗科鲁PalakolluBarony{}

func init() {
    f := BPalakollu波罗科鲁.(*波罗科鲁PalakolluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palakollu",
		TitleName: "波罗科鲁",
		TitleCode: "b_palakollu",
	}
}
