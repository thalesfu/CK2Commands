package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨特迦拉SatgharaBarony struct {
	feud.BaseBarony
}

var BSatghara萨特迦拉 feud.Barony = &萨特迦拉SatgharaBarony{}

func init() {
	f := BSatghara萨特迦拉.(*萨特迦拉SatgharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satghara",
		TitleName: "萨特迦拉",
		TitleCode: "b_satghara",
	}
}
