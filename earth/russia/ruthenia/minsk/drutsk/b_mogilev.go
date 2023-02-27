package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫吉廖夫MogilevBarony struct {
	feud.BaseBarony
}

var BMogilev莫吉廖夫 feud.Barony = &莫吉廖夫MogilevBarony{}

func init() {
    f := BMogilev莫吉廖夫.(*莫吉廖夫MogilevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mogilev",
		TitleName: "莫吉廖夫",
		TitleCode: "b_mogilev",
	}
}
