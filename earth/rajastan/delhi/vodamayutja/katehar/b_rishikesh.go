package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 颉利史计舍RishikeshBarony struct {
	feud.BaseBarony
}

var BRishikesh颉利史计舍 feud.Barony = &颉利史计舍RishikeshBarony{}

func init() {
	f := BRishikesh颉利史计舍.(*颉利史计舍RishikeshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rishikesh",
		TitleName: "颉利史计舍",
		TitleCode: "b_rishikesh",
	}
}
