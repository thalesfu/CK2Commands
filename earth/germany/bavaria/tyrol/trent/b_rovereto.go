package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗韦雷托RoveretoBarony struct {
	feud.BaseBarony
}

var BRovereto罗韦雷托 feud.Barony = &罗韦雷托RoveretoBarony{}

func init() {
	f := BRovereto罗韦雷托.(*罗韦雷托RoveretoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rovereto",
		TitleName: "罗韦雷托",
		TitleCode: "b_rovereto",
	}
}
