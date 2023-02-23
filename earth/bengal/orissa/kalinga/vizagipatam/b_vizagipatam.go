package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗舍佉波吒南VizagipatamBarony struct {
	feud.BaseBarony
}

var BVizagipatam毗舍佉波吒南 feud.Barony = &毗舍佉波吒南VizagipatamBarony{}

func init() {
	f := BVizagipatam毗舍佉波吒南.(*毗舍佉波吒南VizagipatamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vizagipatam",
		TitleName: "毗舍佉波吒南",
		TitleCode: "b_vizagipatam",
	}
}
