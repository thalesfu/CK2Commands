package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡迭石QadeshBarony struct {
	feud.BaseBarony
}

var BQadesh卡迭石 feud.Barony = &卡迭石QadeshBarony{}

func init() {
    f := BQadesh卡迭石.(*卡迭石QadeshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qadesh",
		TitleName: "卡迭石",
		TitleCode: "b_qadesh",
	}
}
