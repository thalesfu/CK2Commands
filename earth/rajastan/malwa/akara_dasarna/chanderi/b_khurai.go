package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库赖KhuraiBarony struct {
	feud.BaseBarony
}

var BKhurai库赖 feud.Barony = &库赖KhuraiBarony{}

func init() {
    f := BKhurai库赖.(*库赖KhuraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khurai",
		TitleName: "库赖",
		TitleCode: "b_khurai",
	}
}
