package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内迈特新堡NemetujvarBarony struct {
	feud.BaseBarony
}

var BNemetujvar内迈特新堡 feud.Barony = &内迈特新堡NemetujvarBarony{}

func init() {
    f := BNemetujvar内迈特新堡.(*内迈特新堡NemetujvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemetujvar",
		TitleName: "内迈特新堡",
		TitleCode: "b_nemetujvar",
	}
}
