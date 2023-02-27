package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德布腊利巴诺斯DebrelibanosBarony struct {
	feud.BaseBarony
}

var BDebrelibanos德布腊利巴诺斯 feud.Barony = &德布腊利巴诺斯DebrelibanosBarony{}

func init() {
    f := BDebrelibanos德布腊利巴诺斯.(*德布腊利巴诺斯DebrelibanosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debrelibanos",
		TitleName: "德布腊利巴诺斯",
		TitleCode: "b_debrelibanos",
	}
}
