package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安布瓦兹AmboiseBarony struct {
	feud.BaseBarony
}

var BAmboise安布瓦兹 feud.Barony = &安布瓦兹AmboiseBarony{}

func init() {
    f := BAmboise安布瓦兹.(*安布瓦兹AmboiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amboise",
		TitleName: "安布瓦兹",
		TitleCode: "b_amboise",
	}
}
