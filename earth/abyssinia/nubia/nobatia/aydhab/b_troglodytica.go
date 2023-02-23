package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗葛罗蒂提卡TroglodyticaBarony struct {
	feud.BaseBarony
}

var BTroglodytica特罗葛罗蒂提卡 feud.Barony = &特罗葛罗蒂提卡TroglodyticaBarony{}

func init() {
	f := BTroglodytica特罗葛罗蒂提卡.(*特罗葛罗蒂提卡TroglodyticaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troglodytica",
		TitleName: "特罗葛罗蒂提卡",
		TitleCode: "b_troglodytica",
	}
}
