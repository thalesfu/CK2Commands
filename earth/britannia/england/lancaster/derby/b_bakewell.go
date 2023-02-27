package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝克韦尔BakewellBarony struct {
	feud.BaseBarony
}

var BBakewell贝克韦尔 feud.Barony = &贝克韦尔BakewellBarony{}

func init() {
    f := BBakewell贝克韦尔.(*贝克韦尔BakewellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakewell",
		TitleName: "贝克韦尔",
		TitleCode: "b_bakewell",
	}
}
