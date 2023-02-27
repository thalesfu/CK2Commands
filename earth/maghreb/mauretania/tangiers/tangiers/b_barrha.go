package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉哈BarrhaBarony struct {
	feud.BaseBarony
}

var BBarrha巴拉哈 feud.Barony = &巴拉哈BarrhaBarony{}

func init() {
    f := BBarrha巴拉哈.(*巴拉哈BarrhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barrha",
		TitleName: "巴拉哈",
		TitleCode: "b_barrha",
	}
}
