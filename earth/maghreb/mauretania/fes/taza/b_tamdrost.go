package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔姆德鲁斯特TamdrostBarony struct {
	feud.BaseBarony
}

var BTamdrost塔姆德鲁斯特 feud.Barony = &塔姆德鲁斯特TamdrostBarony{}

func init() {
	f := BTamdrost塔姆德鲁斯特.(*塔姆德鲁斯特TamdrostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamdrost",
		TitleName: "塔姆德鲁斯特",
		TitleCode: "b_tamdrost",
	}
}
