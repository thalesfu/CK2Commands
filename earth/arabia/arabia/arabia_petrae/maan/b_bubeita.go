package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 部贝塔BubeitaBarony struct {
	feud.BaseBarony
}

var BBubeita部贝塔 feud.Barony = &部贝塔BubeitaBarony{}

func init() {
    f := BBubeita部贝塔.(*部贝塔BubeitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bubeita",
		TitleName: "部贝塔",
		TitleCode: "b_bubeita",
	}
}
