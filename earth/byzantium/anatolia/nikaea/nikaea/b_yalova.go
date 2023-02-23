package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚洛瓦YalovaBarony struct {
	feud.BaseBarony
}

var BYalova亚洛瓦 feud.Barony = &亚洛瓦YalovaBarony{}

func init() {
	f := BYalova亚洛瓦.(*亚洛瓦YalovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yalova",
		TitleName: "亚洛瓦",
		TitleCode: "b_yalova",
	}
}
