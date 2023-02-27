package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥沙OshaBarony struct {
	feud.BaseBarony
}

var BOsha奥沙 feud.Barony = &奥沙OshaBarony{}

func init() {
    f := BOsha奥沙.(*奥沙OshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osha",
		TitleName: "奥沙",
		TitleCode: "b_osha",
	}
}
