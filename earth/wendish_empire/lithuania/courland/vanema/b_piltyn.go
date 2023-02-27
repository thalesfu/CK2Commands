package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔滕PiltynBarony struct {
	feud.BaseBarony
}

var BPiltyn皮尔滕 feud.Barony = &皮尔滕PiltynBarony{}

func init() {
    f := BPiltyn皮尔滕.(*皮尔滕PiltynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piltyn",
		TitleName: "皮尔滕",
		TitleCode: "b_piltyn",
	}
}
