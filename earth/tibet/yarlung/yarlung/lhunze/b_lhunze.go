package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆子LhunzeBarony struct {
	feud.BaseBarony
}

var BLhunze隆子 feud.Barony = &隆子LhunzeBarony{}

func init() {
	f := BLhunze隆子.(*隆子LhunzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhunze",
		TitleName: "隆子",
		TitleCode: "b_lhunze",
	}
}
