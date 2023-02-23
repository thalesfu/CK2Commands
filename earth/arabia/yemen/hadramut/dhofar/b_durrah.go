package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜拉赫DurrahBarony struct {
	feud.BaseBarony
}

var BDurrah杜拉赫 feud.Barony = &杜拉赫DurrahBarony{}

func init() {
	f := BDurrah杜拉赫.(*杜拉赫DurrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durrah",
		TitleName: "杜拉赫",
		TitleCode: "b_durrah",
	}
}
