package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿彭策尔AppenzellBarony struct {
	feud.BaseBarony
}

var BAppenzell阿彭策尔 feud.Barony = &阿彭策尔AppenzellBarony{}

func init() {
    f := BAppenzell阿彭策尔.(*阿彭策尔AppenzellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "appenzell",
		TitleName: "阿彭策尔",
		TitleCode: "b_appenzell",
	}
}
