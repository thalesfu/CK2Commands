package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔东Kordon_vagayBarony struct {
	feud.BaseBarony
}

var BKordon_vagay科尔东 feud.Barony = &科尔东Kordon_vagayBarony{}

func init() {
    f := BKordon_vagay科尔东.(*科尔东Kordon_vagayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kordon_vagay",
		TitleName: "科尔东",
		TitleCode: "b_kordon_vagay",
	}
}
