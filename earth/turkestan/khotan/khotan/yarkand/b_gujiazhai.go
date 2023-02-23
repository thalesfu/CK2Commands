package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顾家寨GujiazhaiBarony struct {
	feud.BaseBarony
}

var BGujiazhai顾家寨 feud.Barony = &顾家寨GujiazhaiBarony{}

func init() {
	f := BGujiazhai顾家寨.(*顾家寨GujiazhaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gujiazhai",
		TitleName: "顾家寨",
		TitleCode: "b_gujiazhai",
	}
}
