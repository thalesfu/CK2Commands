package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷普托TreptowBarony struct {
	feud.BaseBarony
}

var BTreptow特雷普托 feud.Barony = &特雷普托TreptowBarony{}

func init() {
    f := BTreptow特雷普托.(*特雷普托TreptowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treptow",
		TitleName: "特雷普托",
		TitleCode: "b_treptow",
	}
}
