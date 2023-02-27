package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博讷瓦勒BonnevalBarony struct {
	feud.BaseBarony
}

var BBonneval博讷瓦勒 feud.Barony = &博讷瓦勒BonnevalBarony{}

func init() {
    f := BBonneval博讷瓦勒.(*博讷瓦勒BonnevalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonneval",
		TitleName: "博讷瓦勒",
		TitleCode: "b_bonneval",
	}
}
