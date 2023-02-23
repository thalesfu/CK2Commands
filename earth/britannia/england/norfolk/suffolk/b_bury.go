package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝里BuryBarony struct {
	feud.BaseBarony
}

var BBury贝里 feud.Barony = &贝里BuryBarony{}

func init() {
	f := BBury贝里.(*贝里BuryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bury",
		TitleName: "贝里",
		TitleCode: "b_bury",
	}
}
