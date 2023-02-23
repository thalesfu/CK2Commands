package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯摩耆厘HemagiriBarony struct {
	feud.BaseBarony
}

var BHemagiri醯摩耆厘 feud.Barony = &醯摩耆厘HemagiriBarony{}

func init() {
	f := BHemagiri醯摩耆厘.(*醯摩耆厘HemagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hemagiri",
		TitleName: "醯摩耆厘",
		TitleCode: "b_hemagiri",
	}
}
