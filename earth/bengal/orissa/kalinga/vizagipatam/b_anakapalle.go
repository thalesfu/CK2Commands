package vizagipatam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿那迦波罗勒AnakapalleBarony struct {
	feud.BaseBarony
}

var BAnakapalle阿那迦波罗勒 feud.Barony = &阿那迦波罗勒AnakapalleBarony{}

func init() {
    f := BAnakapalle阿那迦波罗勒.(*阿那迦波罗勒AnakapalleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anakapalle",
		TitleName: "阿那迦波罗勒",
		TitleCode: "b_anakapalle",
	}
}
