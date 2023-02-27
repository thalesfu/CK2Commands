package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南旃俱NamchangBarony struct {
	feud.BaseBarony
}

var BNamchang南旃俱 feud.Barony = &南旃俱NamchangBarony{}

func init() {
    f := BNamchang南旃俱.(*南旃俱NamchangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namchang",
		TitleName: "南旃俱",
		TitleCode: "b_namchang",
	}
}
