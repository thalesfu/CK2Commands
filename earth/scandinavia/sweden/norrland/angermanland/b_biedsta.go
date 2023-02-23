package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别德斯塔BiedstaBarony struct {
	feud.BaseBarony
}

var BBiedsta别德斯塔 feud.Barony = &别德斯塔BiedstaBarony{}

func init() {
	f := BBiedsta别德斯塔.(*别德斯塔BiedstaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biedsta",
		TitleName: "别德斯塔",
		TitleCode: "b_biedsta",
	}
}
