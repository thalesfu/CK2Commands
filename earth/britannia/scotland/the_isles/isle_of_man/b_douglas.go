package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道格拉斯DouglasBarony struct {
	feud.BaseBarony
}

var BDouglas道格拉斯 feud.Barony = &道格拉斯DouglasBarony{}

func init() {
    f := BDouglas道格拉斯.(*道格拉斯DouglasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douglas",
		TitleName: "道格拉斯",
		TitleCode: "b_douglas",
	}
}
