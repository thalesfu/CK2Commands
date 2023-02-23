package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉达尼KladanjBarony struct {
	feud.BaseBarony
}

var BKladanj克拉达尼 feud.Barony = &克拉达尼KladanjBarony{}

func init() {
	f := BKladanj克拉达尼.(*克拉达尼KladanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kladanj",
		TitleName: "克拉达尼",
		TitleCode: "b_kladanj",
	}
}
