package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗尔KhorolBarony struct {
	feud.BaseBarony
}

var BKhorol霍罗尔 feud.Barony = &霍罗尔KhorolBarony{}

func init() {
    f := BKhorol霍罗尔.(*霍罗尔KhorolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorol",
		TitleName: "霍罗尔",
		TitleCode: "b_khorol",
	}
}
