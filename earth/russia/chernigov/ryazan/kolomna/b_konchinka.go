package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔钦卡KonchinkaBarony struct {
	feud.BaseBarony
}

var BKonchinka孔钦卡 feud.Barony = &孔钦卡KonchinkaBarony{}

func init() {
	f := BKonchinka孔钦卡.(*孔钦卡KonchinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konchinka",
		TitleName: "孔钦卡",
		TitleCode: "b_konchinka",
	}
}
