package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔德拉TadlaBarony struct {
	feud.BaseBarony
}

var BTadla塔德拉 feud.Barony = &塔德拉TadlaBarony{}

func init() {
	f := BTadla塔德拉.(*塔德拉TadlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadla",
		TitleName: "塔德拉",
		TitleCode: "b_tadla",
	}
}
