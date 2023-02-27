package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普迪维鲁PudiviruBarony struct {
	feud.BaseBarony
}

var BPudiviru普迪维鲁 feud.Barony = &普迪维鲁PudiviruBarony{}

func init() {
    f := BPudiviru普迪维鲁.(*普迪维鲁PudiviruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pudiviru",
		TitleName: "普迪维鲁",
		TitleCode: "b_pudiviru",
	}
}
