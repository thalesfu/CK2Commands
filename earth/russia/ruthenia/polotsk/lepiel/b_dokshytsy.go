package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多克希齐DokshytsyBarony struct {
	feud.BaseBarony
}

var BDokshytsy多克希齐 feud.Barony = &多克希齐DokshytsyBarony{}

func init() {
    f := BDokshytsy多克希齐.(*多克希齐DokshytsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dokshytsy",
		TitleName: "多克希齐",
		TitleCode: "b_dokshytsy",
	}
}
