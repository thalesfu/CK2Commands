package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法捷日FatezhBarony struct {
	feud.BaseBarony
}

var BFatezh法捷日 feud.Barony = &法捷日FatezhBarony{}

func init() {
    f := BFatezh法捷日.(*法捷日FatezhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fatezh",
		TitleName: "法捷日",
		TitleCode: "b_fatezh",
	}
}
