package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔加莫BergamoBarony struct {
	feud.BaseBarony
}

var BBergamo贝尔加莫 feud.Barony = &贝尔加莫BergamoBarony{}

func init() {
    f := BBergamo贝尔加莫.(*贝尔加莫BergamoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergamo",
		TitleName: "贝尔加莫",
		TitleCode: "b_bergamo",
	}
}
