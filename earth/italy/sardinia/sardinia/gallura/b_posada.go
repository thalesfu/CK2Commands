package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波萨达PosadaBarony struct {
	feud.BaseBarony
}

var BPosada波萨达 feud.Barony = &波萨达PosadaBarony{}

func init() {
	f := BPosada波萨达.(*波萨达PosadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posada",
		TitleName: "波萨达",
		TitleCode: "b_posada",
	}
}
