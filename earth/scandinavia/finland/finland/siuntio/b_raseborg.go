package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉塞博里RaseborgBarony struct {
	feud.BaseBarony
}

var BRaseborg拉塞博里 feud.Barony = &拉塞博里RaseborgBarony{}

func init() {
	f := BRaseborg拉塞博里.(*拉塞博里RaseborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raseborg",
		TitleName: "拉塞博里",
		TitleCode: "b_raseborg",
	}
}
