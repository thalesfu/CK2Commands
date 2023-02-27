package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞奇赫SalqahBarony struct {
	feud.BaseBarony
}

var BSalqah塞奇赫 feud.Barony = &塞奇赫SalqahBarony{}

func init() {
    f := BSalqah塞奇赫.(*塞奇赫SalqahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salqah",
		TitleName: "塞奇赫",
		TitleCode: "b_salqah",
	}
}
