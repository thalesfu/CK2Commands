package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里克豪威尔CrickhowellBarony struct {
	feud.BaseBarony
}

var BCrickhowell克里克豪威尔 feud.Barony = &克里克豪威尔CrickhowellBarony{}

func init() {
	f := BCrickhowell克里克豪威尔.(*克里克豪威尔CrickhowellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crickhowell",
		TitleName: "克里克豪威尔",
		TitleCode: "b_crickhowell",
	}
}
