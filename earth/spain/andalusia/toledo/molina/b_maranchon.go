package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马兰琼MaranchonBarony struct {
	feud.BaseBarony
}

var BMaranchon马兰琼 feud.Barony = &马兰琼MaranchonBarony{}

func init() {
    f := BMaranchon马兰琼.(*马兰琼MaranchonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maranchon",
		TitleName: "马兰琼",
		TitleCode: "b_maranchon",
	}
}
