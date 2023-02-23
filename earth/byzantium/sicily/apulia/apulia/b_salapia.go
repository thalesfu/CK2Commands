package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉皮亚SalapiaBarony struct {
	feud.BaseBarony
}

var BSalapia萨拉皮亚 feud.Barony = &萨拉皮亚SalapiaBarony{}

func init() {
	f := BSalapia萨拉皮亚.(*萨拉皮亚SalapiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salapia",
		TitleName: "萨拉皮亚",
		TitleCode: "b_salapia",
	}
}
