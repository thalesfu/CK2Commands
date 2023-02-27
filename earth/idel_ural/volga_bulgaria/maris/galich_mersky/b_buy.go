package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伊BuyBarony struct {
	feud.BaseBarony
}

var BBuy布伊 feud.Barony = &布伊BuyBarony{}

func init() {
    f := BBuy布伊.(*布伊BuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buy",
		TitleName: "布伊",
		TitleCode: "b_buy",
	}
}
