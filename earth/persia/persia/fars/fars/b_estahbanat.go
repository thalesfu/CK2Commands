package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯塔赫巴纳特EstahbanatBarony struct {
	feud.BaseBarony
}

var BEstahbanat埃斯塔赫巴纳特 feud.Barony = &埃斯塔赫巴纳特EstahbanatBarony{}

func init() {
    f := BEstahbanat埃斯塔赫巴纳特.(*埃斯塔赫巴纳特EstahbanatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estahbanat",
		TitleName: "埃斯塔赫巴纳特",
		TitleCode: "b_estahbanat",
	}
}
