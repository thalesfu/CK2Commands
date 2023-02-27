package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗德摩斯ProdromosBarony struct {
	feud.BaseBarony
}

var BProdromos普罗德摩斯 feud.Barony = &普罗德摩斯ProdromosBarony{}

func init() {
    f := BProdromos普罗德摩斯.(*普罗德摩斯ProdromosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prodromos",
		TitleName: "普罗德摩斯",
		TitleCode: "b_prodromos",
	}
}
