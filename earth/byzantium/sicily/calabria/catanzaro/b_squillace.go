package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯奎拉切SquillaceBarony struct {
	feud.BaseBarony
}

var BSquillace斯奎拉切 feud.Barony = &斯奎拉切SquillaceBarony{}

func init() {
	f := BSquillace斯奎拉切.(*斯奎拉切SquillaceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "squillace",
		TitleName: "斯奎拉切",
		TitleCode: "b_squillace",
	}
}
