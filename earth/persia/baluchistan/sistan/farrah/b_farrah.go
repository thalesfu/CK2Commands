package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉FarrahBarony struct {
	feud.BaseBarony
}

var BFarrah法拉 feud.Barony = &法拉FarrahBarony{}

func init() {
	f := BFarrah法拉.(*法拉FarrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farrah",
		TitleName: "法拉",
		TitleCode: "b_farrah",
	}
}
