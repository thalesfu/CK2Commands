package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷焦埃米利亚ReggionellemilaBarony struct {
	feud.BaseBarony
}

var BReggionellemila雷焦埃米利亚 feud.Barony = &雷焦埃米利亚ReggionellemilaBarony{}

func init() {
    f := BReggionellemila雷焦埃米利亚.(*雷焦埃米利亚ReggionellemilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reggionellemila",
		TitleName: "雷焦埃米利亚",
		TitleCode: "b_reggionellemila",
	}
}
