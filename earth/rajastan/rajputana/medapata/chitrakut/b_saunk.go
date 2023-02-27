package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍恩克SaunkBarony struct {
	feud.BaseBarony
}

var BSaunk绍恩克 feud.Barony = &绍恩克SaunkBarony{}

func init() {
    f := BSaunk绍恩克.(*绍恩克SaunkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saunk",
		TitleName: "绍恩克",
		TitleCode: "b_saunk",
	}
}
