package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔萨鲁PiirsaluBarony struct {
	feud.BaseBarony
}

var BPiirsalu皮尔萨鲁 feud.Barony = &皮尔萨鲁PiirsaluBarony{}

func init() {
    f := BPiirsalu皮尔萨鲁.(*皮尔萨鲁PiirsaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piirsalu",
		TitleName: "皮尔萨鲁",
		TitleCode: "b_piirsalu",
	}
}
