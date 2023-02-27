package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈西卜NasibBarony struct {
	feud.BaseBarony
}

var BNasib奈西卜 feud.Barony = &奈西卜NasibBarony{}

func init() {
    f := BNasib奈西卜.(*奈西卜NasibBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasib",
		TitleName: "奈西卜",
		TitleCode: "b_nasib",
	}
}
