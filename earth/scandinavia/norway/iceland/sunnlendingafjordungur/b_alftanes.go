package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔塔内斯AlftanesBarony struct {
	feud.BaseBarony
}

var BAlftanes奥尔塔内斯 feud.Barony = &奥尔塔内斯AlftanesBarony{}

func init() {
	f := BAlftanes奥尔塔内斯.(*奥尔塔内斯AlftanesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alftanes",
		TitleName: "奥尔塔内斯",
		TitleCode: "b_alftanes",
	}
}
