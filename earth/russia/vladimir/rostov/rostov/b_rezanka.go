package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列赞卡RezankaBarony struct {
	feud.BaseBarony
}

var BRezanka列赞卡 feud.Barony = &列赞卡RezankaBarony{}

func init() {
	f := BRezanka列赞卡.(*列赞卡RezankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rezanka",
		TitleName: "列赞卡",
		TitleCode: "b_rezanka",
	}
}
