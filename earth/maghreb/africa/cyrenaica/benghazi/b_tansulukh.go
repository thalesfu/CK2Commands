package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦苏鲁赫TansulukhBarony struct {
	feud.BaseBarony
}

var BTansulukh坦苏鲁赫 feud.Barony = &坦苏鲁赫TansulukhBarony{}

func init() {
    f := BTansulukh坦苏鲁赫.(*坦苏鲁赫TansulukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tansulukh",
		TitleName: "坦苏鲁赫",
		TitleCode: "b_tansulukh",
	}
}
