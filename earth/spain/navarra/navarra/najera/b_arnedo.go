package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔内多ArnedoBarony struct {
	feud.BaseBarony
}

var BArnedo阿尔内多 feud.Barony = &阿尔内多ArnedoBarony{}

func init() {
    f := BArnedo阿尔内多.(*阿尔内多ArnedoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arnedo",
		TitleName: "阿尔内多",
		TitleCode: "b_arnedo",
	}
}
