package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅赫伦MechelenBarony struct {
	feud.BaseBarony
}

var BMechelen梅赫伦 feud.Barony = &梅赫伦MechelenBarony{}

func init() {
    f := BMechelen梅赫伦.(*梅赫伦MechelenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mechelen",
		TitleName: "梅赫伦",
		TitleCode: "b_mechelen",
	}
}
