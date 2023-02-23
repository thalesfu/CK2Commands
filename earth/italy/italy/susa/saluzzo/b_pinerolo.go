package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮内罗洛PineroloBarony struct {
	feud.BaseBarony
}

var BPinerolo皮内罗洛 feud.Barony = &皮内罗洛PineroloBarony{}

func init() {
	f := BPinerolo皮内罗洛.(*皮内罗洛PineroloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinerolo",
		TitleName: "皮内罗洛",
		TitleCode: "b_pinerolo",
	}
}
