package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苗韦尼MioveniBarony struct {
	feud.BaseBarony
}

var BMioveni苗韦尼 feud.Barony = &苗韦尼MioveniBarony{}

func init() {
    f := BMioveni苗韦尼.(*苗韦尼MioveniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mioveni",
		TitleName: "苗韦尼",
		TitleCode: "b_mioveni",
	}
}
