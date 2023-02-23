package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘俱KhuigBarony struct {
	feud.BaseBarony
}

var BKhuig丘俱 feud.Barony = &丘俱KhuigBarony{}

func init() {
	f := BKhuig丘俱.(*丘俱KhuigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khuig",
		TitleName: "丘俱",
		TitleCode: "b_khuig",
	}
}
