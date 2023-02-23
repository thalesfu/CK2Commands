package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安法InfaBarony struct {
	feud.BaseBarony
}

var BInfa安法 feud.Barony = &安法InfaBarony{}

func init() {
	f := BInfa安法.(*安法InfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "infa",
		TitleName: "安法",
		TitleCode: "b_infa",
	}
}
