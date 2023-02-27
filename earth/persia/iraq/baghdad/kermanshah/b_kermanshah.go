package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔曼沙赫KermanshahBarony struct {
	feud.BaseBarony
}

var BKermanshah克尔曼沙赫 feud.Barony = &克尔曼沙赫KermanshahBarony{}

func init() {
    f := BKermanshah克尔曼沙赫.(*克尔曼沙赫KermanshahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kermanshah",
		TitleName: "克尔曼沙赫",
		TitleCode: "b_kermanshah",
	}
}
