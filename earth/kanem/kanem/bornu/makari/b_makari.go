package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马卡里MakariBarony struct {
	feud.BaseBarony
}

var BMakari马卡里 feud.Barony = &马卡里MakariBarony{}

func init() {
	f := BMakari马卡里.(*马卡里MakariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makari",
		TitleName: "马卡里",
		TitleCode: "b_makari",
	}
}
