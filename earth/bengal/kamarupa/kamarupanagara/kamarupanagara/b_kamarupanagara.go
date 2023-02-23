package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦摩缕波城KamarupanagaraBarony struct {
	feud.BaseBarony
}

var BKamarupanagara迦摩缕波城 feud.Barony = &迦摩缕波城KamarupanagaraBarony{}

func init() {
	f := BKamarupanagara迦摩缕波城.(*迦摩缕波城KamarupanagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamarupanagara",
		TitleName: "迦摩缕波城",
		TitleCode: "b_kamarupanagara",
	}
}
