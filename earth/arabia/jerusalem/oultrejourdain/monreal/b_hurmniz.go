package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍姆尼HurmnizBarony struct {
	feud.BaseBarony
}

var BHurmniz霍姆尼 feud.Barony = &霍姆尼HurmnizBarony{}

func init() {
	f := BHurmniz霍姆尼.(*霍姆尼HurmnizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hurmniz",
		TitleName: "霍姆尼",
		TitleCode: "b_hurmniz",
	}
}
