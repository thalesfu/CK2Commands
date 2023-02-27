package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃苏勒VesoulBarony struct {
	feud.BaseBarony
}

var BVesoul沃苏勒 feud.Barony = &沃苏勒VesoulBarony{}

func init() {
    f := BVesoul沃苏勒.(*沃苏勒VesoulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vesoul",
		TitleName: "沃苏勒",
		TitleCode: "b_vesoul",
	}
}
