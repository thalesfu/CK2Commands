package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔塔莫PaltamoBarony struct {
	feud.BaseBarony
}

var BPaltamo帕尔塔莫 feud.Barony = &帕尔塔莫PaltamoBarony{}

func init() {
    f := BPaltamo帕尔塔莫.(*帕尔塔莫PaltamoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paltamo",
		TitleName: "帕尔塔莫",
		TitleCode: "b_paltamo",
	}
}
