package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔MeirBarony struct {
	feud.BaseBarony
}

var BMeir米尔 feud.Barony = &米尔MeirBarony{}

func init() {
	f := BMeir米尔.(*米尔MeirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meir",
		TitleName: "米尔",
		TitleCode: "b_meir",
	}
}
