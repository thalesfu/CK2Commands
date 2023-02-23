package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波伊科夫斯基PoykovskiyBarony struct {
	feud.BaseBarony
}

var BPoykovskiy波伊科夫斯基 feud.Barony = &波伊科夫斯基PoykovskiyBarony{}

func init() {
	f := BPoykovskiy波伊科夫斯基.(*波伊科夫斯基PoykovskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poykovskiy",
		TitleName: "波伊科夫斯基",
		TitleCode: "b_poykovskiy",
	}
}
