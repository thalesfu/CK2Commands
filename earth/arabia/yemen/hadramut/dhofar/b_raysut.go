package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖苏特RaysutBarony struct {
	feud.BaseBarony
}

var BRaysut赖苏特 feud.Barony = &赖苏特RaysutBarony{}

func init() {
	f := BRaysut赖苏特.(*赖苏特RaysutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raysut",
		TitleName: "赖苏特",
		TitleCode: "b_raysut",
	}
}
