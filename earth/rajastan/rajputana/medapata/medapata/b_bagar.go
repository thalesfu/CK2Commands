package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆迦尔BagarBarony struct {
	feud.BaseBarony
}

var BBagar婆迦尔 feud.Barony = &婆迦尔BagarBarony{}

func init() {
    f := BBagar婆迦尔.(*婆迦尔BagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagar",
		TitleName: "婆迦尔",
		TitleCode: "b_bagar",
	}
}
