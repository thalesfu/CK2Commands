package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫利纳MolinaBarony struct {
	feud.BaseBarony
}

var BMolina莫利纳 feud.Barony = &莫利纳MolinaBarony{}

func init() {
	f := BMolina莫利纳.(*莫利纳MolinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molina",
		TitleName: "莫利纳",
		TitleCode: "b_molina",
	}
}
