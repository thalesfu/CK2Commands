package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔瓦ParwaBarony struct {
	feud.BaseBarony
}

var BParwa帕尔瓦 feud.Barony = &帕尔瓦ParwaBarony{}

func init() {
    f := BParwa帕尔瓦.(*帕尔瓦ParwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parwa",
		TitleName: "帕尔瓦",
		TitleCode: "b_parwa",
	}
}
