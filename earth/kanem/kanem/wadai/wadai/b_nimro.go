package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼姆罗NimroBarony struct {
	feud.BaseBarony
}

var BNimro尼姆罗 feud.Barony = &尼姆罗NimroBarony{}

func init() {
	f := BNimro尼姆罗.(*尼姆罗NimroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimro",
		TitleName: "尼姆罗",
		TitleCode: "b_nimro",
	}
}
