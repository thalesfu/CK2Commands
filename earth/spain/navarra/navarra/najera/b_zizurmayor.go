package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹祖拉梅拉ZizurmayorBarony struct {
	feud.BaseBarony
}

var BZizurmayor兹祖拉梅拉 feud.Barony = &兹祖拉梅拉ZizurmayorBarony{}

func init() {
	f := BZizurmayor兹祖拉梅拉.(*兹祖拉梅拉ZizurmayorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zizurmayor",
		TitleName: "兹祖拉梅拉",
		TitleCode: "b_zizurmayor",
	}
}
