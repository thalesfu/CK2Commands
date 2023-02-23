package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗约德拉卢斯ArroyodelalluzBarony struct {
	feud.BaseBarony
}

var BArroyodelalluz阿罗约德拉卢斯 feud.Barony = &阿罗约德拉卢斯ArroyodelalluzBarony{}

func init() {
	f := BArroyodelalluz阿罗约德拉卢斯.(*阿罗约德拉卢斯ArroyodelalluzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arroyodelalluz",
		TitleName: "阿罗约德拉卢斯",
		TitleCode: "b_arroyodelalluz",
	}
}
