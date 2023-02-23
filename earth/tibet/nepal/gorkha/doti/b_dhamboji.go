package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昙菩耆DhambojiBarony struct {
	feud.BaseBarony
}

var BDhamboji昙菩耆 feud.Barony = &昙菩耆DhambojiBarony{}

func init() {
	f := BDhamboji昙菩耆.(*昙菩耆DhambojiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhamboji",
		TitleName: "昙菩耆",
		TitleCode: "b_dhamboji",
	}
}
