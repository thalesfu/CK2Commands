package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉萨尔萨LazarzaBarony struct {
	feud.BaseBarony
}

var BLazarza拉萨尔萨 feud.Barony = &拉萨尔萨LazarzaBarony{}

func init() {
    f := BLazarza拉萨尔萨.(*拉萨尔萨LazarzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lazarza",
		TitleName: "拉萨尔萨",
		TitleCode: "b_lazarza",
	}
}
