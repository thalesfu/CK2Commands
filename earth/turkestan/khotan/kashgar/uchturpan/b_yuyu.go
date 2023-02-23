package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郁于YuyuBarony struct {
	feud.BaseBarony
}

var BYuyu郁于 feud.Barony = &郁于YuyuBarony{}

func init() {
	f := BYuyu郁于.(*郁于YuyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuyu",
		TitleName: "郁于",
		TitleCode: "b_yuyu",
	}
}
