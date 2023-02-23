package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔费尔斯海姆TaversheimBarony struct {
	feud.BaseBarony
}

var BTaversheim塔费尔斯海姆 feud.Barony = &塔费尔斯海姆TaversheimBarony{}

func init() {
	f := BTaversheim塔费尔斯海姆.(*塔费尔斯海姆TaversheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taversheim",
		TitleName: "塔费尔斯海姆",
		TitleCode: "b_taversheim",
	}
}
