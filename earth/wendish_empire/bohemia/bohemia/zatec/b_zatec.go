package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎泰茨ZatecBarony struct {
	feud.BaseBarony
}

var BZatec扎泰茨 feud.Barony = &扎泰茨ZatecBarony{}

func init() {
    f := BZatec扎泰茨.(*扎泰茨ZatecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zatec",
		TitleName: "扎泰茨",
		TitleCode: "b_zatec",
	}
}
