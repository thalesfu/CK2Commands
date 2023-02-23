package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙特法拉维克NattfaravikBarony struct {
	feud.BaseBarony
}

var BNattfaravik瑙特法拉维克 feud.Barony = &瑙特法拉维克NattfaravikBarony{}

func init() {
	f := BNattfaravik瑙特法拉维克.(*瑙特法拉维克NattfaravikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nattfaravik",
		TitleName: "瑙特法拉维克",
		TitleCode: "b_nattfaravik",
	}
}
