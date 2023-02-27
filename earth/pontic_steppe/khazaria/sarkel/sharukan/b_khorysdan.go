package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍雷斯丹KhorysdanBarony struct {
	feud.BaseBarony
}

var BKhorysdan霍雷斯丹 feud.Barony = &霍雷斯丹KhorysdanBarony{}

func init() {
    f := BKhorysdan霍雷斯丹.(*霍雷斯丹KhorysdanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorysdan",
		TitleName: "霍雷斯丹",
		TitleCode: "b_khorysdan",
	}
}
