package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佛尔萨ForsaBarony struct {
	feud.BaseBarony
}

var BForsa佛尔萨 feud.Barony = &佛尔萨ForsaBarony{}

func init() {
	f := BForsa佛尔萨.(*佛尔萨ForsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forsa",
		TitleName: "佛尔萨",
		TitleCode: "b_forsa",
	}
}
