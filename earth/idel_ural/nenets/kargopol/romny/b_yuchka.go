package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤奇卡YuchkaBarony struct {
	feud.BaseBarony
}

var BYuchka尤奇卡 feud.Barony = &尤奇卡YuchkaBarony{}

func init() {
    f := BYuchka尤奇卡.(*尤奇卡YuchkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuchka",
		TitleName: "尤奇卡",
		TitleCode: "b_yuchka",
	}
}
