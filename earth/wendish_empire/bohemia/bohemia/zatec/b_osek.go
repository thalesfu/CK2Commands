package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥塞克OsekBarony struct {
	feud.BaseBarony
}

var BOsek奥塞克 feud.Barony = &奥塞克OsekBarony{}

func init() {
    f := BOsek奥塞克.(*奥塞克OsekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osek",
		TitleName: "奥塞克",
		TitleCode: "b_osek",
	}
}
