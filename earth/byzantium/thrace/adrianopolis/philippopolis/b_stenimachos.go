package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯忒尼马科斯StenimachosBarony struct {
	feud.BaseBarony
}

var BStenimachos斯忒尼马科斯 feud.Barony = &斯忒尼马科斯StenimachosBarony{}

func init() {
    f := BStenimachos斯忒尼马科斯.(*斯忒尼马科斯StenimachosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stenimachos",
		TitleName: "斯忒尼马科斯",
		TitleCode: "b_stenimachos",
	}
}
