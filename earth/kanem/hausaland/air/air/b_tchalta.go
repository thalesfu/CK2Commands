package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查尔塔TchaltaBarony struct {
	feud.BaseBarony
}

var BTchalta查尔塔 feud.Barony = &查尔塔TchaltaBarony{}

func init() {
    f := BTchalta查尔塔.(*查尔塔TchaltaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tchalta",
		TitleName: "查尔塔",
		TitleCode: "b_tchalta",
	}
}
