package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈登GordonBarony struct {
	feud.BaseBarony
}

var BGordon戈登 feud.Barony = &戈登GordonBarony{}

func init() {
	f := BGordon戈登.(*戈登GordonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gordon",
		TitleName: "戈登",
		TitleCode: "b_gordon",
	}
}
