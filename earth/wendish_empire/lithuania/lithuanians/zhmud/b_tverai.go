package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特韦赖TveraiBarony struct {
	feud.BaseBarony
}

var BTverai特韦赖 feud.Barony = &特韦赖TveraiBarony{}

func init() {
    f := BTverai特韦赖.(*特韦赖TveraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tverai",
		TitleName: "特韦赖",
		TitleCode: "b_tverai",
	}
}
