package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆拉恩LampronBarony struct {
	feud.BaseBarony
}

var BLampron拉姆拉恩 feud.Barony = &拉姆拉恩LampronBarony{}

func init() {
    f := BLampron拉姆拉恩.(*拉姆拉恩LampronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lampron",
		TitleName: "拉姆拉恩",
		TitleCode: "b_lampron",
	}
}
