package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰多尔RanderBarony struct {
	feud.BaseBarony
}

var BRander兰多尔 feud.Barony = &兰多尔RanderBarony{}

func init() {
    f := BRander兰多尔.(*兰多尔RanderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rander",
		TitleName: "兰多尔",
		TitleCode: "b_rander",
	}
}
