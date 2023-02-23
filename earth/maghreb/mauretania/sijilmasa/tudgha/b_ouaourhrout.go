package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦乌尔鲁OuaourhroutBarony struct {
	feud.BaseBarony
}

var BOuaourhrout瓦乌尔鲁 feud.Barony = &瓦乌尔鲁OuaourhroutBarony{}

func init() {
	f := BOuaourhrout瓦乌尔鲁.(*瓦乌尔鲁OuaourhroutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouaourhrout",
		TitleName: "瓦乌尔鲁",
		TitleCode: "b_ouaourhrout",
	}
}
