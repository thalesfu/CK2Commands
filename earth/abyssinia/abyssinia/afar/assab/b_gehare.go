package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格哈尔GehareBarony struct {
	feud.BaseBarony
}

var BGehare格哈尔 feud.Barony = &格哈尔GehareBarony{}

func init() {
    f := BGehare格哈尔.(*格哈尔GehareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gehare",
		TitleName: "格哈尔",
		TitleCode: "b_gehare",
	}
}
