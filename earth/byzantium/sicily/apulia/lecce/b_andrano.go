package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德拉诺AndranoBarony struct {
	feud.BaseBarony
}

var BAndrano安德拉诺 feud.Barony = &安德拉诺AndranoBarony{}

func init() {
    f := BAndrano安德拉诺.(*安德拉诺AndranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andrano",
		TitleName: "安德拉诺",
		TitleCode: "b_andrano",
	}
}
