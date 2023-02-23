package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肖奈ChonaiBarony struct {
	feud.BaseBarony
}

var BChonai肖奈 feud.Barony = &肖奈ChonaiBarony{}

func init() {
	f := BChonai肖奈.(*肖奈ChonaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chonai",
		TitleName: "肖奈",
		TitleCode: "b_chonai",
	}
}
