package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔赫姆格拉VerkhmglaBarony struct {
	feud.BaseBarony
}

var BVerkhmgla韦尔赫姆格拉 feud.Barony = &韦尔赫姆格拉VerkhmglaBarony{}

func init() {
    f := BVerkhmgla韦尔赫姆格拉.(*韦尔赫姆格拉VerkhmglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verkhmgla",
		TitleName: "韦尔赫姆格拉",
		TitleCode: "b_verkhmgla",
	}
}
