package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达罗摩尔格波利DayrodmorgabrielBarony struct {
	feud.BaseBarony
}

var BDayrodmorgabriel达罗摩尔格波利 feud.Barony = &达罗摩尔格波利DayrodmorgabrielBarony{}

func init() {
	f := BDayrodmorgabriel达罗摩尔格波利.(*达罗摩尔格波利DayrodmorgabrielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dayrodmorgabriel",
		TitleName: "达罗摩尔格波利",
		TitleCode: "b_dayrodmorgabriel",
	}
}
