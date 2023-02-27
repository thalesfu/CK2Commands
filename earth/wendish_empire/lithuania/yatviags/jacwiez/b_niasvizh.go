package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅斯维日NiasvizhBarony struct {
	feud.BaseBarony
}

var BNiasvizh涅斯维日 feud.Barony = &涅斯维日NiasvizhBarony{}

func init() {
    f := BNiasvizh涅斯维日.(*涅斯维日NiasvizhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niasvizh",
		TitleName: "涅斯维日",
		TitleCode: "b_niasvizh",
	}
}
