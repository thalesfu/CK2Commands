package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗洛拉克FloracBarony struct {
	feud.BaseBarony
}

var BFlorac弗洛拉克 feud.Barony = &弗洛拉克FloracBarony{}

func init() {
    f := BFlorac弗洛拉克.(*弗洛拉克FloracBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "florac",
		TitleName: "弗洛拉克",
		TitleCode: "b_florac",
	}
}
