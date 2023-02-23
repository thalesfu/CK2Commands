package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩耆厘RamagiriBarony struct {
	feud.BaseBarony
}

var BRamagiri罗摩耆厘 feud.Barony = &罗摩耆厘RamagiriBarony{}

func init() {
	f := BRamagiri罗摩耆厘.(*罗摩耆厘RamagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramagiri",
		TitleName: "罗摩耆厘",
		TitleCode: "b_ramagiri",
	}
}
