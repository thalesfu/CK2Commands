package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡佩拉什KapilashBarony struct {
	feud.BaseBarony
}

var BKapilash卡佩拉什 feud.Barony = &卡佩拉什KapilashBarony{}

func init() {
    f := BKapilash卡佩拉什.(*卡佩拉什KapilashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapilash",
		TitleName: "卡佩拉什",
		TitleCode: "b_kapilash",
	}
}
