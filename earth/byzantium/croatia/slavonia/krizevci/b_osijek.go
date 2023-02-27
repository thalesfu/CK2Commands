package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥西耶克OsijekBarony struct {
	feud.BaseBarony
}

var BOsijek奥西耶克 feud.Barony = &奥西耶克OsijekBarony{}

func init() {
    f := BOsijek奥西耶克.(*奥西耶克OsijekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osijek",
		TitleName: "奥西耶克",
		TitleCode: "b_osijek",
	}
}
